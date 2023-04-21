package user

import (
	"github.com/labstack/echo/v4"
)

var _ Handler = &user{}

type Handler interface {
	Create() echo.HandlerFunc
}

type user struct {
}

func NewUserHandler() Handler {
	return &user{}
}
