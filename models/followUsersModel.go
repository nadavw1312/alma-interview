package models

type FollowUser struct {
	Id         string `json:"id"`
	FollowerId string `json:"followerId"`
	FollowedId string `json:"followedId"`
}

type NewFollowUser struct {
	FollowerId string `json:"followerId"`
	FollowedId string `json:"followedId"`
}
