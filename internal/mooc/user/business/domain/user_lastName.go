package domain

import "github.com/erik-sostenes/users-api/internal/shared/domain"

type UserLastName struct {
	value string
}

func NewUserLastName(value string) (UserLastName, error) {
	accountLastName, err := domain.String(value).Validate()
	if err != nil {
		return UserLastName{}, err
	}

	return UserLastName{
		value: accountLastName,
	}, nil
}

func (a UserLastName) String() string {
	return a.value
}
