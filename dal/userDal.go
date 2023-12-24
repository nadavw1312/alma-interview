package dal

import (
	"context"

	"github.com/nadavw1312/golang-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDal struct {
	collection *mongo.Collection
}

func NewUserDal(collection *mongo.Collection) *UserDal {
	return &UserDal{collection}
}

func (ud *UserDal) FindByUsername(ctx context.Context, userName string) (models.User, error) {
	return BaseFindOneByQuery[models.User](ctx, ud.collection, bson.M{"username": userName})
}

func (userDal *UserDal) GetById(ctx context.Context, id string) (models.User, error) {
	return BaseFindById[models.User](ctx, userDal.collection, id)
}

func (userDal *UserDal) InsertUser(ctx context.Context, user models.CreateUserRequest) (string, error) {
	return BaseInsert(ctx, userDal.collection, user)
}

func (userDal *UserDal) DeleteById(ctx context.Context, id string) error {
	return BaseDeleteById(ctx, userDal.collection, id)
}
