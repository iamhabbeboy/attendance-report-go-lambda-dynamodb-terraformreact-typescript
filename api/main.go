package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	database := NewDatabaseDriver(DefaultDatabaseDriver)
	DB = database

	const (
		FirstService = iota + 1
		SecondService
	)

	serv := SecondService
	fmt.Println(serv)

	r := NewRoutes().Endpoints()

	if os.Getenv("ENV") == "development" {
		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
	} else {
		// lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
		lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			response := events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Success",
			}
			return response, nil
		})
	}
}
