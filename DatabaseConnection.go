package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DatabaseConnection struct {
	access     bool              // access to access database
	uri        string            // connection string of mongodb
	client     *mongo.Client     // mongodb client
	collection *mongo.Collection // mongodb collection
}

// Connecting to Database
func (d *DatabaseConnection) connect() {
	if d.access {
		// Connect to database
		d.uri = os.Getenv("MONGODB_URI")
		client, err := mongo.Connect(options.Client().ApplyURI(d.uri))
		if err != nil {
			panic(err)
		}
		d.client = client
		d.collection = d.client.Database("webCrawlerArchive").Collection("webpages")
		filter := bson.D{{}}
		// Delete all documents in the collection from previous crawl
		d.collection.DeleteMany(context.TODO(), filter)
	}
}

// disconnecting database
func (d *DatabaseConnection) disconnect() {
	if d.access {
		d.client.Disconnect(context.TODO())
	}
}

// insert webpage into the database
func (d *DatabaseConnection) insertWebpage(webpage Webpage) {
	if d.access {
		d.collection.InsertOne(context.TODO(), webpage)
	}
}
