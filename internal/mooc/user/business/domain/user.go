package domain

type User struct {
	UserId
	UserName
	UserLastName
}

func NewUser(userId UserId, userName UserName, userLastName UserLastName) (User, error) {
	return User{
		userId,
		userName,
		userLastName,
	}, nil
}
