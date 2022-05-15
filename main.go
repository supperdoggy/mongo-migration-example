package main

import (
	"context"
	"log"
	"time"

	_ "github.com/supperdoggy/migration_test/migrations"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	/*
	   Connect to my cluster
	*/
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("test")
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		log.Fatal(err)
	}

	log.Println("success")

}
