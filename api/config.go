package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	GetMethod             = "GET"
	PostMethod            = "POST"
	DefaultDBName         = "attendance"
	DefaultDatabaseDriver = "dynamoDb"
)

var DB interface{}

func JSON(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, err = w.Write([]byte("error occured with data"))
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		return
	}
}
