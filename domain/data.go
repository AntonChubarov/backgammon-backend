package domain

import "fmt"

var UserExistError = fmt.Errorf("user with this login already exists")
var InvalidLogin = fmt.Errorf("user with this login is not registered")
var InvalidPassword = fmt.Errorf("invalid password")
