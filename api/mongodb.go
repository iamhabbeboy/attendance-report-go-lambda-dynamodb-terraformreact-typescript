package main

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodber struct {
	col *mongo.Collection
}

func NewMongo() *Mongodber {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	coll := client.Database(DefaultDBName).Collection(DefaultTableName)

	return &Mongodber{
		col: coll,
	}
}

func GetID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func (db *Mongodber) Store(report Report) (interface{}, error) {
	result, err := db.col.InsertOne(context.TODO(), report)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	if result.InsertedID == nil {
		errMsg := "Error occured when inserting record"
		log.Printf(errMsg)
		return primitive.ObjectID{}, errors.New(errMsg)
	}
	return result.InsertedID, nil
}

func (a *Mongodber) Get(ID string) error {
	return nil
}
