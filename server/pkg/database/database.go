package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/config"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDB() {

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority",
		config.Cnfg.DBUser, config.Cnfg.DBPassword, config.Cnfg.DBHost)

	// Create a context with a 10-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB.

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check the connection.
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("Failed to ping MongoDB.", err)
	}
	db = client.Database(config.Cnfg.DBName)
	logger.Info("Database Connect successfully.")
}

func GetCollection(ctx *gin.Context, collectionName string) *mongo.Collection {
	collection := db.Collection(collectionName)

	return collection
}

func Save(ctx *gin.Context, data interface{}, collectionName string) error {
	collection := db.Collection(collectionName)
	_, err := collection.InsertOne(ctx, data)
	return err
}

func SaveMany(ctx *gin.Context, data []interface{}, collectionName string) error {
	collection := db.Collection(collectionName)
	_, err := collection.InsertMany(ctx, data)
	return err
}

func Fetch(ctx *gin.Context, collectionName string) ([]bson.M, error) {
	collection := db.Collection(collectionName)

	// Find documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var documents []bson.M
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return documents, nil
}

func FetchByKeyValue(ctx *gin.Context, collectionName string, key string, value interface{}) ([]bson.M, error) {
	collection := db.Collection(collectionName)

	// Define filter based on key and value
	filter := bson.M{key: value}

	// Find documents in the collection by key-value pair
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var documents []bson.M
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return documents, nil
}
