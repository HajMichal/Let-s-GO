package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {

	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("There is no .env file", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("There is no mongodb connection", err)
	}
	fmt.Println("Connected to mongoDB")
	log.Fatal(app.Listen(os.Getenv("PORT")))
}