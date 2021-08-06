package dal

type UserDBDTO struct {
	Login string `db:"login"`
	PasswordHash string `db:"password"`
}