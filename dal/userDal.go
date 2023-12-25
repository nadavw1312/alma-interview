package dal

import (
	"context"
	"errors"

	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/utils"
)

type UserDal struct {
	Users map[string]*models.User
}

func NewUserDal() *UserDal {
	return &UserDal{Users: make(map[string]*models.User)}
}

func (userDal *UserDal) GetById(id string) (models.User, error) {
	if _, ok := userDal.Users[id]; ok {
		return *userDal.Users[id], nil
	}

	return models.User{}, errors.New("User not found")
}

func (userDal *UserDal) GetAll() ([]*models.User, error) {
	users := []*models.User{}
	for _, user := range userDal.Users {
		users = append(users, user)
	}
	return users, nil
}

func (userDal *UserDal) InsertUser(ctx context.Context, user models.CreateUserRequest) (string, error) {
	for _, existingUser := range userDal.Users {
		if existingUser.Name == user.Name {
			return "", errors.New("User already exists")
		}
	}

	newUser := &models.User{Name: user.Name}
	id, err := utils.GenerateRandomID()
	if err != nil {
		return "", err
	}
	newUser.Id = id
	userDal.Users[id] = newUser
	return id, nil
}

func (userDal *UserDal) DeleteById(ctx context.Context, id string) error {
	if _, ok := userDal.Users[id]; ok {
		delete(userDal.Users, id)
		return nil
	}

	return errors.New("User not found")
}

func (userDal *UserDal) UpdateById(ctx context.Context, id string, user models.User) error {
	if _, ok := userDal.Users[id]; ok {
		userDal.Users[id].Name = user.Name
		return nil
	}

	return errors.New("User not found")
}
