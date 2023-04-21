package main

import "github.com/erik-sostenes/users-api/internal/apps/backend/dependency"

func main() {
	if err := dependency.NewInjector(); err != nil {
		panic(err)
	}
}
