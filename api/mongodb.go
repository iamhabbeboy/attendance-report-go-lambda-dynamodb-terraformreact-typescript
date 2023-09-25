package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(DefaultDBName).Collection(DefaultTableName)

	return &Mongodber{
		col: coll,
	}
}

func (db *Mongodber) Store(report Report) error {
	result, err := db.col.InsertOne(context.TODO(), report)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return nil
}

func (a *Mongodber) Get(ID string) error {
	return nil
}
