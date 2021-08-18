package auth

import (
	"backgammon/config"
	domainAuth "backgammon/domain/auth"
	"backgammon/utils"
	"log"
)

type UserAuthService struct {
	storage            domainAuth.UserDataStorage
	mainSessionStorage SessionStorage
	config             *config.ServerConfig
	hasher             StringHasher
	tokenGenerator     TokenGenerator
}

func NewUserAuthService(storage domainAuth.UserDataStorage,
	mainSessionStorage SessionStorage,
	config *config.ServerConfig,
	tokenGenerator TokenGenerator) *UserAuthService {
	return &UserAuthService{storage: storage,
		mainSessionStorage: mainSessionStorage,
		config:             config,
		hasher:             NewHasherSHA256(),
		tokenGenerator:     tokenGenerator}
}

func (uas *UserAuthService) RegisterNewUser(data domainAuth.UserAuthData) error {
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
