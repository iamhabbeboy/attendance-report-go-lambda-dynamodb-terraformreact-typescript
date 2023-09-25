package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpResponse struct {
}

func String(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}

func JSON(w http.ResponseWriter, data interface{}, code int) {
	w.WriteHeader(code)
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
