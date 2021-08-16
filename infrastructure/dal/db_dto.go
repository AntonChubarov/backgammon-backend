package dal

type UserDBDTO struct {
	UUID  string `db:"useruuid"`
	Login string `db:"userlogin"`
	PasswordHash string `db:"userpassword"`
}