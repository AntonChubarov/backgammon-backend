package auth

import "fmt"

var ErrorMoreThanOneUsernameRecord = fmt.Errorf("more than one user with this username, report to developers was automatically send")
var ErrorNoUserInDatabase = fmt.Errorf("no user with this username in database")
