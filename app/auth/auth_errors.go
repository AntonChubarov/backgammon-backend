package auth

import "fmt"

var ErrorUserExists = fmt.Errorf("user with this username already exists")
var ErrorInvalidUsername = fmt.Errorf("user with this username is not registered")
var ErrorInvalidPassword = fmt.Errorf("invalid password")
