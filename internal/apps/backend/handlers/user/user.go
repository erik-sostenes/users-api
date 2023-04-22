package user

import (
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/services/create"
	"github.com/erik-sostenes/users-api/internal/shared/domain/bus/command"
	"github.com/labstack/echo/v4"
)

var _ Handler = &user{}

type Handler interface {
	Create() echo.HandlerFunc
}

type user struct {
	command.Bus[create.UserCommand]
}

func NewUserHandler(command command.Bus[create.UserCommand]) Handler {
	return &user{
		command,
	}
}
