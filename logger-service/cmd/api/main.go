package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort = "9999"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	loadConfig()
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start web server
	// go app.serve()
	log.Println("Starting service on port", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}

}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", Cnfg.DBUser, Cnfg.DBPassword, Cnfg.DBHost)
	fmt.Println(connStr)
	// Create a context with a 10-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return client, nil
}
