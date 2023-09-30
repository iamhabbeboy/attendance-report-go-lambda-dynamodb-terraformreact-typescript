package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoDbConfig struct {
	table dynamo.Table
}

func NewDynamoDB(tableName string) *DynamoDbConfig {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if tableName == "" {
		log.Fatal("Table name must not be empty")
	}
	return &DynamoDbConfig{
		table: db.Table(tableName),
	}
}

func (d *DynamoDbConfig) Store(report Report) (interface{}, error) {
	err := d.table.Put(dynamo.AWSEncoding(report)).Run()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (d *DynamoDbConfig) Get(ID string) error {
	var result Report
	err := d.table.Get("id", ID).
		One(&result)
	if err != nil {
		return err
	}
	return nil
}
