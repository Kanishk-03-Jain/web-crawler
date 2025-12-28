package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DatabaseConnection struct {
	access     bool
	uri        string
	client     *mongo.Client
	collection *mongo.Collection
}

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
		// Delete all documents in the collection
		d.collection.DeleteMany(context.TODO(), filter)
	}
}

func (d *DatabaseConnection) disconnect() {
	if d.access {
		d.client.Disconnect(context.TODO())
	}
}

func (d *DatabaseConnection) insertWebpage(webpage Webpage) {
	if d.access {
		d.collection.InsertOne(context.TODO(), webpage)
	}
}
