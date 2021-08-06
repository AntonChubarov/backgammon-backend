package domain

import "fmt"

var UserExistError = fmt.Errorf("user with this login already exists")
