package types

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nadavw1312/golang-fiber/env"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(env.Env.AppSecret)

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
