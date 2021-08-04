package httphandlers

import "github.com/labstack/echo/v4"

type UserRegistrator struct {
	
}

func NewUserRegistrator() *UserRegistrator {
	return &UserRegistrator{}
}

func (ur *UserRegistrator) Handle(c echo.Context) error {
	return nil
}