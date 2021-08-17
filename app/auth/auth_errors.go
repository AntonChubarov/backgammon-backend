package auth

import "fmt"

var ErrorUserExists = fmt.Errorf("user with this login already exists")
var ErrorInvalidLogin = fmt.Errorf("user with this login is not registered")
var ErrorInvalidPassword = fmt.Errorf("invalid password")
