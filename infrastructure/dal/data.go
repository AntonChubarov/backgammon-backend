package dal

import "fmt"

var MoreThanOneLoginRecordError = fmt.Errorf("more than one user with this login, report to developers was automatically send")
