package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/config"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/util/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	logger.Info("Database Connect successfully.")
}
