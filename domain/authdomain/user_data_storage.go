package authdomain

type UserStorage interface {
	AddNewUser(*UserData) error
	GetUserByUsername(UserName) (UserData, error)
	GetUserByUUID(UUID) (UserData, error)
	UpdateUser(UUID, *UserData) error
	RemoveUser(UUID) error
}


