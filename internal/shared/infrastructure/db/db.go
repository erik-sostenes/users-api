package db

import "github.com/erik-sostenes/users-api/internal/shared"

// Type represents an uint for the type of DataBase
type Type uint

const (
	// SQL represents MySQL database
	SQL Type = iota
	// NoSQL represents MongoDB database
	NoSQL
)

// Configuration represents the settings of the type of storage
type Configuration struct {
	// Type defines the type of storage to be used.
	Type
	Driver   string
	Host     string
	Port     string
	User     string
	Database string
	Password string
}

// NewMongoDBConfiguration returns an instance of Configuration with all the settings
// to make the connection to the database
func NewMongoDBConfiguration() Configuration {
	return Configuration{
		Type:     NoSQL,
		Driver:   shared.GetEnv("NoSQL_DRIVER"),
		Host:     shared.GetEnv("NoSQL_HOST"),
		Port:     shared.GetEnv("NoSQL_PORT"),
		Database: shared.GetEnv("NoSQL_DATABASE"),
	}
}
