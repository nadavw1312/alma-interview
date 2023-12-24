package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nadavw1312/golang-fiber/utils"
)

type ENV struct {
	MongoURI  string `validate:"required"`
	DbName    string `validate:"required"`
	AppSecret string `validate:"required"`
}

var Env = ENV{}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	Env.DbName = os.Getenv("DB_NAME")
	Env.MongoURI = os.Getenv("MONGO_URI")
	Env.AppSecret = os.Getenv("APP_SECRET")

	if errs := utils.Validator.Validate(Env); len(errs.Errors) > 0 {
		message := ""
		for _, v := range errs.Errors {
			message += v.Message + "\n"
		}
		log.Fatal("Env variables are missing\n" + message)
	}

}

func GetEnv(key string) string {
	return os.Getenv(key)
}
