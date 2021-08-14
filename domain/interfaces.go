package domain

type UserStorage interface {
	AddNewUser(data UserData) error
	IsUserExist(login string) (bool, error)
	GetUserByLogin(login string) (UserData, error)
	UpdateUser(oldData UserData, newData UserData) error
	RemoveUser(data UserData) error
}
