package data

import "go.mongodb.org/mongo-driver/mongo"

type DataLayer struct {
	collection *mongo.Collection
}

func InitDataLayer(collection *mongo.Collection) *DataLayer {
	return &DataLayer{
		collection: collection,
	}
}
