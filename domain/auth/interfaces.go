package auth

type UserDataStorage interface {
	AddNewUser(data UserAuthData) error
	IsUserExist(username string) (bool, error)
	GetUserByUsername(username string) (UserAuthData, error)
	UpdateUser(oldData UserAuthData, newData UserAuthData) error
	RemoveUser(data UserAuthData) error
}


