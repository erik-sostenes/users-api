package domain

import "context"

type (
	Command interface {
		CommandId() string
	}

	UserCreator[V Command] interface {
		Create(context.Context, V) error
	}
)
