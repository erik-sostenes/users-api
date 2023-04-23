package create

import (
	"context"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
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
	userId, err := domain.NewUserId(cmd.Id)
	if err != nil {
		return err
	}

	userName, err := domain.NewUserName(cmd.Name)
	if err != nil {
		return err
	}

	userLastName, err := domain.NewUserLastName(cmd.LastName)
	if err != nil {
		return err
	}
	return c.UserCreator.Create(ctx, userId, userName, userLastName)
}
