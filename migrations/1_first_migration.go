package migrations

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		ctx := context.Background()
		_, err := db.Collection("students").UpdateMany(ctx, bson.M{"name": "old guy"}, bson.M{"$set": bson.M{"age": 45}})
		return err
	}, func(db *mongo.Database) error {
		ctx := context.Background()
		_, err := db.Collection("students").UpdateMany(ctx, bson.M{"name": "old guy"}, bson.M{"$set": bson.M{"age": 44}})
		return err
	})
}
