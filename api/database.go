package main

import "log"

type Databaser interface {
	Store(report Report) error
	Get(ID string) error
	// Delete()
}

func NewDatabaseDriver(driver string) Databaser {
	if driver == "" {
		log.Fatal("Database service cannot be empty")
	}
	srvs := map[string]Databaser{
		"dynamoDb":  NewDynamoDB(DefaultDBName),
		"firestore": NewFirestore(),
		"mongodb":   NewMongo(),
	}
	return srvs[driver]
}
