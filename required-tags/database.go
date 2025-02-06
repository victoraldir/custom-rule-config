package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "TagsTable"

type DynamoDBClient interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

type DynamoDBRepository interface {
	PutItem(objectCatalog CatalogObject) (*dynamodb.PutItemOutput, error)
	GetItem(objectId string) (*CatalogObject, error)
}

type DynamoDBRepositoryImpl struct {
	client DynamoDBClient
}

// NewDynamoDBRepository creates a new instance of DynamoDBRepositoryImpl
func NewDynamoDBRepository(sess *session.Session) *DynamoDBRepositoryImpl {
	return &DynamoDBRepositoryImpl{client: dynamodb.New(sess)}
}

// PutItem inserts a CatalogObject into the DynamoDB table
func (s *DynamoDBRepositoryImpl) PutItem(objectCatalog CatalogObject) (*dynamodb.PutItemOutput, error) {
	input := &dynamodb.PutItemInput{
		Item:      objectCatalog.ToItem(),
		TableName: aws.String(tableName),
	}

	return s.client.PutItem(input)
}

// GetItem retrieves a CatalogObject from the DynamoDB table by objectId
func (s *DynamoDBRepositoryImpl) GetItem(objectId string) (*CatalogObject, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ObjectId": {
				S: aws.String(objectId),
			},
		},
		TableName: aws.String(tableName),
	}

	item, err := s.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	if len(item.Item) == 0 {
		return nil, nil
	}

	objectCatalog := CatalogObject{
		ObjectId: objectId,
		Tags:     make(map[string]string),
	}

	if len(item.Item) == 0 {
		return &objectCatalog, nil
	}

	for key, value := range item.Item["Tags"].M {
		objectCatalog.Tags[key] = *value.S
	}

	return &objectCatalog, nil
}
