package authdomain

type UUID string
type Token string
type UserName string
type Password string


type AutorizedUserData struct {
	UUID
	UserName
	Password
	Token
}

type UserData struct {
	UUID
	UserName
	Password
}

func (u *UserData) Update(sample *UserData) {
	u.UserName=sample.UserName
	u.Password=sample.Password
}

