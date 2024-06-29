package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"lambda/model"
)

type DynamoStorage struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func (storage *DynamoStorage) SavePersona(ctx context.Context, persona *model.Persona) error {
	av, err := dynamodbattribute.MarshalMap(persona)
	fmt.Printf("Save Persona. MarshalMap: %s\n", av)

	if err != nil {
		fmt.Printf("MarshalMap Error: %s \n", err.Error())
		return err
	}
	fmt.Printf("Save Persona. TableName: %s\n", storage.tableName)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &storage.tableName,
	}
	_, err = storage.svc.PutItemWithContext(ctx, input)
	if err != nil {
		fmt.Printf("PutItemWithContext Error: %s\n", err.Error())
	}
	return err
}

func NewDynamoStorage(sess *session.Session, tableName string) *DynamoStorage {
	fmt.Printf("New Dynamo Storage. TableName: %s\n", tableName)
	return &DynamoStorage{
		svc:       dynamodb.New(sess),
		tableName: tableName,
	}
}
