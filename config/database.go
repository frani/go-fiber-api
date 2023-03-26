package configs

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Ctx | @desc: mongo context interface
var Ctx context.Context

// Cancel | @desc: mongo context cancel function
var Cancel context.CancelFunc

// Client | @desc: mongo client struct
var client *mongo.Client

// DB | @desc: mongo database struct
var DB *mongo.Database

// Connect | @desc: connects to mongoDB
func ConnectDB() {
	var err error

	Ctx, Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	client, err = mongo.Connect(Ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	// Connect to mongo database
	DB = client.Database(os.Getenv("MONGO_DB"))
}
