package main

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/supperdoggy/migration_test/migrations"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	/*
	   Connect to my cluster
	*/
	args := os.Args[1:]
	log.Println("args", args)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("test")
	migrate.SetDatabase(db)

	switch args[0] {
	case "up":
		if err := migrate.Up(migrate.AllAvailable); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := migrate.Down(migrate.AllAvailable); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("you need to pick up or down")
	}

	log.Println("success")

}
