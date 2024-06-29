package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"lambda/api"
	"lambda/event"
	"lambda/storage"
	"os"
)

func main() {
	// print log
	fmt.Println("Setting Up Persona Lambda")
	sess := session.Must(session.NewSession())
	fmt.Println("Creating A DynamoDB Table")
	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		err := fmt.Errorf("TABLE_NAME:%s environment variable is not set", tableName)
		fmt.Println(err.Error())
	}
	dynamoStorage := storage.NewDynamoStorage(sess, tableName)
	fmt.Println("Creating Publisher")
	publisher := event.NewBridgePublisher(sess)
	fmt.Println("Setting up Api Layer")
	handler := api.NewHandler(dynamoStorage, publisher)
	fmt.Println("Lambda Starting")
	lambda.Start(handler.HandlerRequest)
}
