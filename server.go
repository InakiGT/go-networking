package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	e := echo.New()

	uri, exists := os.LookupEnv("MONGO_URL")
	if !exists {
		uri = "mongodb://localhost:27017"
	}

	client, err := connectMongo(uri)
	if err != nil {
		log.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		id := uuid.New().String()
		var msg = Message{
			id,
			1,
		}
		collection := client.Database("test").Collection("ids")
		_, err := collection.InsertOne(context.TODO(), msg)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":5050"))
}
