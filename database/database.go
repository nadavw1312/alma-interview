package database

import (
	"context"
	"fmt"
	"time"

	"github.com/nadavw1312/golang-fiber/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MG MongoInstance

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Env.MongoURI))
	db := client.Database(env.Env.DbName)

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Print("Connection Opened to Database")
	}

	MG = MongoInstance{Client: client, Db: db}
	return nil
}
