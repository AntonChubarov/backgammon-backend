package authdomain

type SessionStorage interface {
	AddSession (SessionData) error
	GetSessionByToken(Token) (SessionData, error)
	GetSessionSByUUID(UUID) (SessionData, error)
	DeleteSession(Token) error
	UpdateSession(Token, SessionData) error

}
