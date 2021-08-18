package auth

type SessionStorage interface {
	AddNewUser(data *UserSessionData)
	UpdateTokenExpiryTime(token string)
	DeleteUserByToken(token string)
	GetTokenByUUID(uuid string) (token string, wasFound bool)
}

type StringHasher interface {
	HashString(s string) (string, error)
}

type TokenGenerator interface {
	GenerateToken () string
}

