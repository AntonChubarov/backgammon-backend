package dal

type UserDBDTO struct {
	UUID         string `db:"useruuid"`
	Username     string `db:"userlogin"`
	PasswordHash string `db:"userpassword"`
}