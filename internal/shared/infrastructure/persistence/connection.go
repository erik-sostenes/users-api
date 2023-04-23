package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoClient method that will connect a mongoDB and returns an instance of mongo.Database
func NewMongoClient(config Configuration) (*mongo.Database, error) {
	switch config.Type {
	case NoSQL:
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
			fmt.Sprintf("%s://%s:%s", config.Driver, config.Host, config.Port),
		))
		if err != nil {
			return nil, err
		}

		err = client.Ping(context.TODO(), readpref.Primary())
		if err != nil {
			return nil, err
		}

		db := client.Database(config.Database)

		return db, db.Client().Ping(context.TODO(), readpref.Primary())
	default:
		panic(fmt.Sprintf("%T type is not supported", config.Type))
	}
}

// NewMongoDataBase method that returns an instance of mongo.DataBase
// if an error occurs a panic will be launched
func NewMongoDataBase(config Configuration) (db *mongo.Database) {
	db, err := NewMongoClient(config)
	if err != nil {
		panic(err)
	}
	return
}
