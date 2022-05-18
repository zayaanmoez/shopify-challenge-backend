package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	client *mongo.Client
}

func ConnectDB() *Connection {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://root:rootpassword@shopify-ch.d02zr.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to DB and pinged.")

	connection := new(Connection)
	connection.client = client
	return connection
}

func (connection *Connection) OpenCollection(db string, collection string) *mongo.Collection {
	dbinstance := connection.client.Database(db)
	collectionIns := dbinstance.Collection(collection)
	return collectionIns
}
