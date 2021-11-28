package domain

// TODO use only primitives to describe interfaces

type UserEntry interface {
	RegisterUser(credentials UserCredentials) (err error)
	AuthorizeUser(credentials UserCredentials) (userId string, err error)
	UpdateUser(userId, credentials UserCredentials) (err error)
	DeleteUser(userId string) (err error)
}

type UserStorage interface {
	AddNewUser(UserData) error
	GetUserByUsername(UserName) (UserData, error)
	GetUserByUUID(UUID) (UserData, error)
	UpdateUser(UUID, *UserData) error
	RemoveUser(UUID) error
}

type Sessioner interface {
	NewSession(userId string) (token string)
	CheckSession(token string) (isActive bool)
	GetUserIdByToken(token string) (userId string, err error)
	RefreshSession(token string) (err error)
	CloseSession(token string) (err error)
}

type SessionStorage interface {
	AddSession(SessionData) error
	GetSessionByToken(Token) (SessionData, error)
	GetSessionSByUUID(UUID) (SessionData, error)
	DeleteSession(Token) error
	UpdateSession(Token, SessionData) error
}
