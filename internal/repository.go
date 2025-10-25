package internal

import (
	"context"
	"log"
	"time"

	"github.com/samedozturk/Auth-Service/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func StartMongoDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(config.GetEnv("MONGO_URI")))
	if err != nil {
		log.Fatal("database connection error:", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("database ping error:", err)
	}
	return client.Database(config.GetEnv("MONGODB_NAME"))
}
