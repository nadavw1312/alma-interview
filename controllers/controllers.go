package controllers

var UserControllerP *UserController
var FollowUsersControllerP *FollowUsersController
var TweetControllerP *TweetController

func InitControllers() {
	UserControllerP = NewUserController()
	FollowUsersControllerP = NewFollowUsersController()
	TweetControllerP = NewTweetController()
}
