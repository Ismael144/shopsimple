package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func EnsureProductIndexes(ctx context.Context, db *mongo.Database) error {
	col := db.Collection("products")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "name", Value: "text"}}, 
		}, 
		{
			Keys: bson.D{{Key: "price_minor", Value: 1}}, 
		}, 
		{
			Keys: bson.D{{Key: "categories", Value: 1}},
		}, 
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		}, 
	}

	_, err := col.Indexes().CreateMany(ctx, indexes)
	return err
}