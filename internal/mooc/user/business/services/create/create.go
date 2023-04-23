package create

import (
	"context"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
)

type UserCreator struct {
	domain.UserRepository[domain.UserId, domain.User]
}

func (u *UserCreator) Create(ctx context.Context, userId, userName, userLastName string) (err error) {
	user, err := domain.NewUser(userId, userName, userLastName)
	if err != nil {
		return
	}

	return u.Save(ctx, user.UserId, user)
}
