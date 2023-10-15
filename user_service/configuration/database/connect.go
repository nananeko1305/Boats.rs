package connection

import (
	"context"
	"fmt"
	configuration "user_service/configuration/enviroment"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initClient(enviroment *configuration.Enviroment) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", enviroment.USER_DB_HOST, enviroment.USER_DB_PORT)
	optionsClient := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), optionsClient)
}

func ConnectToDatabase(enviroment *configuration.Enviroment) (*mongo.Collection, error) {
	client, err := initClient(enviroment)
	if err != nil {
		return nil, err
	}
	return client.Database(enviroment.DATABASE).Collection(enviroment.COLLECTION), nil
}
