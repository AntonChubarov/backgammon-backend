package domain

type UserStorage interface {
	AddNewUser(data UserData) error
	IsUserExist(data UserData) (bool, error)
	UpdateUser(oldData UserData, newData UserData) error
	RemoveUser(data UserData) error
}
