package dal

import (
	"errors"

	"github.com/nadavw1312/golang-fiber/models"
	"github.com/nadavw1312/golang-fiber/utils"
)

type FollowUserDal struct {
	followUsers map[string]*models.FollowUser
}

func NewFollowUserDal() *FollowUserDal {
	return &FollowUserDal{followUsers: make(map[string]*models.FollowUser)}
}

func (followUserDal *FollowUserDal) Follow(followerId string, followedId string) (string, error) {
	id, err := utils.GenerateRandomID()
	if err != nil {
		return "", err
	}
	followUser := &models.FollowUser{Id: id, FollowerId: followerId, FollowedId: followedId}
	followUserDal.followUsers[id] = followUser
	return id, nil
}

func (followUserDal *FollowUserDal) Unfollow(followerId string, followedId string) error {
	for _, followUser := range followUserDal.followUsers {
		if followUser.FollowerId == followerId && followUser.FollowedId == followedId {
			delete(followUserDal.followUsers, followUser.Id)
			return nil
		}
	}

	return errors.New("User not found")
}

func (dal *FollowUserDal) GetById(id string) (*models.FollowUser, error) {
	if _, ok := dal.followUsers[id]; ok {
		return dal.followUsers[id], nil
	}

	return nil, errors.New("User not found")
}

func (dal *FollowUserDal) GetByFollowerId(followerId string) ([]*models.FollowUser, error) {
	var followUsers []*models.FollowUser

	for _, followUser := range dal.followUsers {
		if followUser.FollowerId == followerId {
			followUsers = append(followUsers, followUser)
		}
	}

	return followUsers, nil
}
