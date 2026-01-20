package db

import "go.mongodb.org/mongo-driver/v2/mongo"

func NewMongoDatabase(client *mongo.Client, dbName string) *mongo.Database {
	return client.Database(dbName)
}