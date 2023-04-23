package domain

import "github.com/erik-sostenes/users-api/internal/shared/domain"

type UserId struct {
	value string
}

func NewUserId(value string) (UserId, error) {
	id, err := domain.Identifier(value).ParseUuID()
	if err != nil {
		return UserId{}, err
	}

	return UserId{
		value: id,
	}, nil
}

func (u UserId) String() string {
	return u.value
}
