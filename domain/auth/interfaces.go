package auth

type UserDataStorage interface {
	AddNewUser(data UserAuthData) error
	IsUserExist(login string) (bool, error)
	GetUserByLogin(login string) (UserAuthData, error)
	UpdateUser(oldData UserAuthData, newData UserAuthData) error
	RemoveUser(data UserAuthData) error
}

type SessionStorage interface {
	AddNewUser(data UserSessionData)
	UpdateTokenExpiryTime(token string)
	DeleteUserByToken(token string)
	GetTokenByUUID(uuid string) (token string, wasFound bool)
}
