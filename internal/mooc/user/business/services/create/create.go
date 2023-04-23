package create

import (
	"context"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
)

type UserCreator struct {
	domain.UserRepository[domain.UserId, domain.User]
}

func (u *UserCreator) Create(ctx context.Context,
	userId domain.UserId,
	userName domain.UserName,
	userLastName domain.UserLastName) error {

	user, err := domain.NewUser(userId, userName, userLastName)
	if err != nil {
		return err
	}

	return u.Save(ctx, userId, user)
}
