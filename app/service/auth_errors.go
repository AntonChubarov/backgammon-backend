package service

import "fmt"

//Registration Errors
var ErrorUserExists = fmt.Errorf("user with this username already exists")
var ErrorPoorUsername = fmt.Errorf("username don't meet the requirements: it should contain at least six characters, one letter, and one digit")
var ErrorPoorPassword = fmt.Errorf("password don't meet the requirements: it should contain at least eight characters, one letter, and one digit")

//AutorizationErrors
var ErrorUserNotRegistered = fmt.Errorf("user with this username is not registered")
var ErrorInvalidPassword = fmt.Errorf("invalid password")

//SessioningErrors
var ErrorInvalidToken = fmt.Errorf("the token you provide is invalid")
var ErrorNoActiveSessions = fmt.Errorf("user with this uuid has no active sessions")
var ErrorDuplicateSession = fmt.Errorf("this session already stored in session")
var ErrorUserMultiSessioning = fmt.Errorf("this user already has open session")

var ErrorNullArgument = fmt.Errorf("null argument exception")
