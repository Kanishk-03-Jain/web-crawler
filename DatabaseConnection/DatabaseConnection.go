package databaseconnection

import (
	"context"
	"os"

	parsinghelpers "github.com/Kanishk-03-Jain/web-crawler/parsingHelpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
	access     bool
	uri        string
	client     *mongo.Client
	collection *mongo.Collection
}

func (d *DatabaseConnection) Connect() {
	if d.access {
		// Connect to database
		d.uri = os.Getenv("MONGODB_URI")
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.uri))
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

func (d *DatabaseConnection) Disconnect() {
	if d.access {
		d.client.Disconnect(context.TODO())
	}
}

func (d *DatabaseConnection) InsertWebpage(webpage parsinghelpers.Webpage) {
	if d.access {
		d.collection.InsertOne(context.TODO(), webpage)
	}
}
