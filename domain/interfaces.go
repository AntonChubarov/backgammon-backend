package domain

type UserStorage interface {
	AddNewUser(info UserInfo) error
	CheckUser(info UserInfo) error
	UpdateUser(oldInfo UserInfo, newInfo UserInfo) error
	RemoveUser(info UserInfo) error
}
