package database

import (
	"context"
	"fmt"

	"github.com/mertdogan12/cn/internal/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	database   *mongo.Database
	collection *mongo.Collection
)

func Connect() (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?maxPoolSize=20&w=majority", conf.Database.Username, conf.Database.Password, conf.Database.Host)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	fmt.Println("Erfolgreich mit der Mongo Datanbank verbunden")

	database = client.Database(conf.Database.Database)
	collection = database.Collection("rechner")

	return client, nil
}

func Disconnect(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}

	return nil
}

func GetName(id string) (string, error) {
	var result bson.M
	if err := collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result); err != nil {
		return "", err
	}

	data := result["name"]

	return data.(string), nil
}
