package app

import "backgammon/domain"

type UserRegistrationService struct {
	storage domain.UserStorage
}

func NewUserRegistrationService(storage domain.UserStorage) *UserRegistrationService {
	return &UserRegistrationService{storage: storage}
}

func (urs *UserRegistrationService) RegisterNewUser(data domain.UserData) error {
	userExist, err := urs.storage.IsUserExist(data)
	if userExist {
		return domain.UserExistError
	}
	if err != nil {
		return err
	}

	data.Password, err = HashPassword(data.Password)
	if err != nil {
		return err
	}

	err = urs.storage.AddNewUser(data)
	if err != nil {
		return err
	}

	return nil
}