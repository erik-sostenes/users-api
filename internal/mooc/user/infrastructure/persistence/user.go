package persistence

import (
	"context"
	"errors"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	*mongo.Collection
}

func NewUserStore(db *mongo.Collection) domain.UserRepository[domain.UserId, domain.User] {
	return &UserStore{
		db,
	}
}

func (u *UserStore) Save(ctx context.Context, userId domain.UserId, user domain.User) (err error) {
	if _, err = u.InsertOne(ctx, bson.M{
		"_id":       userId.String(),
		"name":      user.UserName.String(),
		"last_name": user.UserLastName.String(),
	}); err != nil {
		if ok := mongo.IsDuplicateKeyError(err); ok {
			return ErrDuplicateUser
		}
		return
	}
	return
}

func (u *UserStore) Delete(ctx context.Context, userId domain.UserId) (err error) {
	if _, err = u.DeleteOne(ctx, bson.D{{"_id", userId.String()}}); err != nil {
		return
	}
	return
}

func (u *UserStore) Find(ctx context.Context, userId domain.UserId) (domain.User, error) {
	filter := bson.M{"_id": userId.String()}

	var data bson.M
	err := u.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, err
		}
	}

	return domain.NewUser(data["_id"].(string), data["name"].(string), data["last_name"].(string))
}
