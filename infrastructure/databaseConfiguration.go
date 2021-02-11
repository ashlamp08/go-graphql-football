package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Mongodb *mongo.Database

func (env *Environment) InitMongoDB() (db *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(env.Database["mongodb"].Connection)
	client, err := mongo.Connect(context.Background(), clientOptions)
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("OH SHIT")
		return db, err
	}
	Mongodb = client.Database(env.Database["mongodb"].Name)
	log.Println("Mongodb is Ready!!")
	return db, err
}
