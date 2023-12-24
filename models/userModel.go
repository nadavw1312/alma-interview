package models

type User struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username" validate:"required"`
	Password string `json:"-" bson:"password" validate:"required"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Gender   string `json:"gender" bson:"gender" validate:"required"`
	Age      int    `json:"age" bson:"age" validate:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username,omitempty" bson:"username" validate:"required"`
	Password string `json:"password,omitempty" bson:"password" validate:"required"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Gender   string `json:"gender" bson:"gender" validate:"required"`
	Age      int    `json:"age" bson:"age" validate:"required"`
}
