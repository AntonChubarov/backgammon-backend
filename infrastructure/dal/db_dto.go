package dal

type UserDBDTO struct {
	Id string `db:"userUUID"`
	Login string `db:"userlogin"`
	PasswordHash string `db:"userpassword"`
}