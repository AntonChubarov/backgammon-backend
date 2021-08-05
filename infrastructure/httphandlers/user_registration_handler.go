package httphandlers

import (
	"backgammon/app"
	"github.com/labstack/echo/v4"
)

type UserRegistrator struct {
	service app.UserRegistrationService
}

func NewUserRegistrator() *UserRegistrator {
	return &UserRegistrator{}
}

func (ur *UserRegistrator) Handle(c echo.Context) error {
	return nil
}