package models

type User struct {
	Id   string `json:"id"`
	Name string `json:"name" bson:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name" bson:"name"`
}
