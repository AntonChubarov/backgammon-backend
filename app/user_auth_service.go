package app

import (
	"backgammon/config"
	"backgammon/domain"
	"log"
)

type UserAuthService struct {
	storage domain.UserDataStorage
	mainSessionStorage domain.SessionStorage
	config *config.ServerConfig
}

func NewUserAuthService(storage domain.UserDataStorage, mainSessionStorage domain.SessionStorage, config *config.ServerConfig) *UserAuthService {
	return &UserAuthService{storage: storage, mainSessionStorage: mainSessionStorage, config: config}
}

func (uas *UserAuthService) RegisterNewUser(data domain.UserAuthData) error {
	userExist, err := uas.storage.IsUserExist(data.Login)
	if userExist {
		return UserExistError
	}

	data.UUID = GenerateUUID()

	data.Password, err = HashPassword(data.Password)
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

func (uas *UserAuthService) AuthorizeUser(data domain.UserAuthData) (token string, err error) {
	token = ""
	var user domain.UserAuthData

	user, err = uas.storage.GetUserByLogin(data.Login)
	if err != nil {
		return
	}

	var passwordHash string
	passwordHash, err = HashPassword(data.Password)
	if err != nil {
		return
	}

	if passwordHash != user.Password {
		err = InvalidPassword
		return
	}

	var wasFound bool
	token, wasFound = uas.mainSessionStorage.GetTokenByUUID(user.UUID)
	if wasFound {
		return
	}
	token = GenerateToken(uas.config.Token.TokenLength, uas.config.Token.TokenSymbols)
	uas.mainSessionStorage.AddNewUser(domain.UserSessionData{Token: token, UserUUID: user.UUID})
	return
}