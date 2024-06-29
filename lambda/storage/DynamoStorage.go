package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
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

func (storage *DynamoStorage) GetPersonaByID(ctx context.Context, id string) (*model.Persona, error) {
	fmt.Printf("Getting Person by ID:: %s\n", id)
	input := &dynamodb.GetItemInput{
		TableName: &storage.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := storage.svc.GetItemWithContext(ctx, input)
	if err != nil {
		fmt.Printf("GetItemWithContext Error: %s\n", err.Error())
		return nil, err
	}

	if result.Item == nil {
		fmt.Printf("Persona with Id: %s cannot be find", id)
		return nil, nil
	}

	var persona model.Persona
	err = dynamodbattribute.UnmarshalMap(result.Item, &persona)
	if err != nil {
		fmt.Printf("UnmarshalMap Error: %s\n", err.Error())
		return nil, err
	}

	return &persona, nil
}

func (storage *DynamoStorage) GetAllPersonas(ctx context.Context) ([]*model.Persona, error) {
	input := &dynamodb.ScanInput{
		TableName: &storage.tableName,
	}

	result, err := storage.svc.ScanWithContext(ctx, input)
	if err != nil {
		fmt.Printf("Cannot scan the Db Error: %s\n", err.Error())
		return nil, err
	}

	var personas []*model.Persona
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &personas)
	if err != nil {
		fmt.Printf("UnmarshalMap Error: %s\n", err.Error())
		return nil, err
	}

	return personas, nil
}
