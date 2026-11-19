package database

import (
	"context"
	"time"

	"github.com/mrmissx/stashbin/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var session *mongo.Database

// ConnectDb opens the Mongo client and verifies connectivity before serving.
// DB_URI accepts both mongodb:// and mongodb+srv:// (Atlas).
func ConnectDb() {
	dbUri := utils.GetEnv("DB_URI", "")
	if dbUri == "" {
		utils.Logger.Fatal("DB_URI is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		utils.Logger.Fatal("Failed to create Mongo client: ", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		utils.Logger.Fatal("Failed to ping Mongo: ", err)
	}

	dbName := utils.GetEnv("DB_NAME", "stashbin")
	session = client.Database(dbName)
	utils.Logger.Info("Successfully connected database")
}

func GetDatabase() *mongo.Database {
	return session
}
