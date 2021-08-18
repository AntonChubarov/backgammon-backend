package auth

import (
	"backgammon/config"
	domainAuth "backgammon/domain/auth"
	"backgammon/utils"
	"github.com/dlclark/regexp2"
	"log"
)

type UserAuthService struct {
	storage            domainAuth.UserDataStorage
	mainSessionStorage SessionStorage
	config             *config.ServerConfig
	hasher             StringHasher
	tokenGenerator     TokenGenerator
	usernameRegexp *regexp2.Regexp
	passwordRegexp *regexp2.Regexp
}

func NewUserAuthService(storage domainAuth.UserDataStorage,
	mainSessionStorage SessionStorage,
	config *config.ServerConfig,
	tokenGenerator TokenGenerator) *UserAuthService {
	return &UserAuthService{storage: storage,
		mainSessionStorage: mainSessionStorage,
		config:             config,
		hasher:             NewHasherSHA256(),
		tokenGenerator:     tokenGenerator,
		usernameRegexp: regexp2.MustCompile("^(?=.{6,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$", 0),
		passwordRegexp: regexp2.MustCompile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$", 0),
	}
}

func (uas *UserAuthService) RegisterNewUser(data domainAuth.UserAuthData) error {
	var isMatch bool

	if isMatch, _ = uas.usernameRegexp.MatchString(data.Username); !isMatch {
		return ErrorPoorUsername
	}
	if isMatch, _ = uas.passwordRegexp.MatchString(data.Password); !isMatch {
		return ErrorPoorPassword
	}

	userExist, err := uas.storage.IsUserExist(data.Username)
	if userExist {
		return ErrorUserExists
	}
	if err != nil {
		return err
	}

	data.UUID = utils.GenerateUUID()

	data.Password, err = uas.hasher.HashString(data.Password)
	if err != nil {
		log.Println("In app.RegisterNewUser", err)
		return err
	}

	err = uas.storage.AddNewUser(data)
	if err != nil {
		log.Println("In app.RegisterNewUser", err)
		return err
	}

	return nil
}

func (uas *UserAuthService) AuthorizeUser(data domainAuth.UserAuthData) (token string, err error) {
	// Need to discuss
	var isMatch bool

	if isMatch, _ = uas.usernameRegexp.MatchString(data.Username); !isMatch {
		return "", ErrorUserNotRegistered
	}
	if isMatch, _ = uas.passwordRegexp.MatchString(data.Password); !isMatch {
		return "", ErrorInvalidPassword
	}

	token = ""
	var user domainAuth.UserAuthData

	user, err = uas.storage.GetUserByUsername(data.Username)
	if err != nil {
		return
	}

	var passwordHash string
	passwordHash, err = uas.hasher.HashString(data.Password)
	if err != nil {
		return
	}

	if passwordHash != user.Password {
		err = ErrorInvalidPassword
		return
	}

	var wasFound bool
	token, wasFound = uas.mainSessionStorage.GetTokenByUUID(user.UUID)
	if wasFound {
		return
	}
	token = uas.tokenGenerator.GenerateToken()
	uas.mainSessionStorage.AddNewUser(&UserSessionData{Token: token, UserUUID: user.UUID})
	return
}
