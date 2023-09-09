package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	r := NewRoutes().Endpoints()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
