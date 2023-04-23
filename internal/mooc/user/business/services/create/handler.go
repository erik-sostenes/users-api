package create

import (
	"context"
	"github.com/erik-sostenes/users-api/internal/shared/domain/bus/command"
)

var _ command.Command = UserCommand{}

type UserCommand struct {
	Id       string
	Name     string
	LastName string
}

func (c UserCommand) CommandId() string {
	return "create_user_command"
}

var _ command.Handler[UserCommand] = &CreateUserCommandHandler{}

type CreateUserCommandHandler struct {
	UserCreator
}

func (c *CreateUserCommandHandler) Handler(ctx context.Context, cmd UserCommand) error {
	return c.UserCreator.Create(ctx, cmd.Id, cmd.Name, cmd.LastName)
}
