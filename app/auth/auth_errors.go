package auth

import "fmt"

var ErrorUserExists = fmt.Errorf("user with this username already exists")
var ErrorUserNotRegistered = fmt.Errorf("user with this username is not registered")
var ErrorInvalidPassword = fmt.Errorf("invalid password")
var ErrorPoorUsername = fmt.Errorf("username don't meet the requirements: it should contain at least six characters, one letter, and one digit")
var ErrorPoorPassword = fmt.Errorf("password don't meet the requirements: it should contain at least eight characters, one letter, and one digit")
var ErrorInvalidToken = fmt.Errorf("the token you provide is invalid")