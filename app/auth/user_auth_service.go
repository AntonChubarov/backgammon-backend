package auth

import (
	"backgammon/config"
	"backgammon/domain/authdomain"
	"backgammon/utils"
	"github.com/dlclark/regexp2"
	"log"
	"time"
)

type UserAuthService struct {
	storage            authdomain.UserStorage
	mainSessionStorage authdomain.SessionStorage
	config             *config.ServerConfig
	hasher             StringHasher
	tokenGenerator     TokenGenerator
	usernameRegexp     *regexp2.Regexp
	passwordRegexp     *regexp2.Regexp
}

func NewUserAuthService(storage authdomain.UserStorage,
	mainSessionStorage authdomain.SessionStorage,
	config *config.ServerConfig,
	tokenGenerator TokenGenerator) *UserAuthService {
	return &UserAuthService{storage: storage,
		mainSessionStorage: mainSessionStorage,
		config:             config,
		hasher:             NewHasherSHA256(),
		tokenGenerator:     tokenGenerator,
		usernameRegexp:     regexp2.MustCompile("^(?=.{6,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$", 0),
		passwordRegexp:     regexp2.MustCompile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$", 0),
	}
}

func (uas *UserAuthService) RegisterNewUser(data authdomain.UserData) error {
	var isMatch bool

	if isMatch, _ = uas.usernameRegexp.MatchString(string(data.UserName)); !isMatch {
		return ErrorPoorUsername
	}
	if isMatch, _ = uas.passwordRegexp.MatchString(string(data.Password)); !isMatch {
		return ErrorPoorPassword
	}

	_, err := uas.storage.GetUserByUsername(data.UserName)
	if err == nil {
		log.Println("Register: user", data.UserName, "already exist")
		return ErrorUserExists
	}

	data.UUID = utils.GenerateUUID()

	passwordHash, err := uas.hasher.HashString(string(data.Password))
	if err != nil {
		//log.Println("In app.RegisterNewUser", err)
		return err
	}
	data.Password = authdomain.Password(passwordHash)

	err = uas.storage.AddNewUser(data)
	if err != nil {
		//log.Println("In app.RegisterNewUser", err)
		return err
	}

	log.Println("Register: user", data.UserName, "registered")
	return nil
}

func (uas *UserAuthService) AuthorizeUser(data authdomain.UserData) (token authdomain.Token, err error) {
	var user authdomain.UserData

	user, err = uas.storage.GetUserByUsername(data.UserName)
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
			session.ExpiryTime = authdomain.ExpiryTime(time.Now().UTC().Add(5 * time.Second))
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

	token = authdomain.Token(uas.tokenGenerator.GenerateToken())
	tokenExpiryTime := authdomain.ExpiryTime(time.Now().UTC().Add(5 * time.Second))
	err = uas.mainSessionStorage.AddSession(authdomain.SessionData{UUID: user.UUID, Token: token, ExpiryTime: tokenExpiryTime})
	log.Println("Authorize: user", data.UserName, "authorized, session created")
	return token, err
}

// Need to be refactored
func (uas *UserAuthService) CheckToken(token authdomain.Token) error {
	_, err := uas.mainSessionStorage.GetSessionByToken(token)
	if err != nil {
		return err
	}
	return nil
}
