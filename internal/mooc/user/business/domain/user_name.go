package domain

import "github.com/erik-sostenes/users-api/internal/shared/domain"

type UserName struct {
	value string
}

func NewUserName(value string) (UserName, error) {
	accountName, err := domain.String(value).Validate()
	if err != nil {
		return UserName{}, err
	}

	return UserName{
		value: accountName,
	}, nil
}

func (a UserName) String() string {
	return a.value
}
