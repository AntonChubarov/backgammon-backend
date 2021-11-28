package service

import (
	"backgammon/domain"
	"backgammon/utils"
	"github.com/dlclark/regexp2"
	"github.com/jbrodriguez/mlog"
	"log"
	"time"
)

type UserEntryService struct {
	storage            domain.UserStorage
	mainSessionStorage domain.SessionStorage
	hasher             StringHasher
	tokenGenerator     TokenGenerator
	usernameRegexp     *regexp2.Regexp
	passwordRegexp     *regexp2.Regexp
}

func NewUserAuthService(storage domain.UserStorage,
	mainSessionStorage domain.SessionStorage,
	tokenGenerator TokenGenerator) *UserEntryService {
	return &UserEntryService{
		storage:            storage,
		mainSessionStorage: mainSessionStorage,
		hasher:             NewHasherSHA256(),
		tokenGenerator:     tokenGenerator,
		usernameRegexp:     regexp2.MustCompile("^(?=.{6,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$", 0),
		passwordRegexp:     regexp2.MustCompile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$", 0),
	}
}

func (uas *UserEntryService) RegisterNewUser(data domain.UserData) error {
	if isMatch, _ := uas.usernameRegexp.MatchString(string(data.UserName)); !isMatch {
		return ErrorPoorUsername
	}
	if isMatch, _ := uas.passwordRegexp.MatchString(string(data.Password)); !isMatch {
		return ErrorPoorPassword
	}

	_, err := uas.storage.GetUserByUsername(data.UserName)
	if err == nil {
		mlog.Warning("user %s already exist", data.UserName)
		return ErrorUserExists
	}

	data.UUID = utils.GenerateUUID()

	passwordHash, err := uas.hasher.HashString(string(data.Password))
	if err != nil {
		return err
	}
	data.Password = domain.Password(passwordHash)

	err = uas.storage.AddNewUser(data)
	if err != nil {
		return err
	}

	mlog.Info("Register: user %s registered", data.UserName)
	return nil
}

func (uas *UserEntryService) AuthorizeUser(data domain.UserData) (token domain.Token, err error) {

	user, err := uas.storage.GetUserByUsername(data.UserName)
	if err != nil {
		log.Println("Authorize: user", data.UserName, "not registered")
		return
	}

	var passwordHash string
	passwordHash, err = uas.hasher.HashString(string(data.Password))
	if err != nil {
		return
	}

	if passwordHash != string(user.Password) {
		log.Println("Authorize: user", data.UserName, "entered invalid password")
		err = ErrorInvalidPassword
		return
	}

	session, err := uas.mainSessionStorage.GetSessionSByUUID(user.UUID)

	if err == nil {
		if time.Time(session.ExpiryTime).After(time.Now().UTC()) {
			//TODO refactor prolong session
			session.ExpiryTime = domain.ExpiryTime(time.Now().UTC().Add(5 * time.Second))
			uas.mainSessionStorage.UpdateSession(session.Token, session)
			log.Println("Authorize: user", data.UserName, "has active session, session prolonged")
			return session.Token, nil
		}
		log.Println("Authorize: user", data.UserName, "has expired session, session deleted")
		err = uas.mainSessionStorage.DeleteSession(session.Token)
		if err != nil {
			return "", err
		}

	}

	token = domain.Token(uas.tokenGenerator.GenerateToken())
	//TODO refactor prolong session
	tokenExpiryTime := domain.ExpiryTime(time.Now().UTC().Add(5 * time.Second))
	err = uas.mainSessionStorage.AddSession(domain.SessionData{UUID: user.UUID, Token: token, ExpiryTime: tokenExpiryTime})
	log.Println("Authorize: user", data.UserName, "authorized, session created")
	return token, err
}

// Need to be refactored
func (uas *UserEntryService) CheckToken(token domain.Token) error {
	_, err := uas.mainSessionStorage.GetSessionByToken(token)
	if err != nil {
		return err
	}
	return nil
}