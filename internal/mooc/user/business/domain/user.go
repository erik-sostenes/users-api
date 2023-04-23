package domain

type User struct {
	UserId
	UserName
	UserLastName
}

func NewUser(id, name, lastName string) (User, error) {
	userId, err := NewUserId(id)
	if err != nil {
		return User{}, err
	}

	userName, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	userLastName, err := NewUserLastName(lastName)
	if err != nil {
		return User{}, err
	}
	return User{
		userId,
		userName,
		userLastName,
	}, nil
}
