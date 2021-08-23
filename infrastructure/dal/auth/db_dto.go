package auth

type UserDBDTO struct {
	UUID         string `db:"useruuid"`
	Username     string `db:"username"`
	PasswordHash string `db:"userpassword"`
}