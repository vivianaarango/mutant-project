package repositories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoAPI interface for DynamoDB methods.
type DynamoAPI interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
}

// DynamoProvider interface for Dynamo client.
type DynamoProvider interface {
	DynamoClient() (DynamoAPI, error)
}

// DynamoDB struct for dynamo client.
type DynamoDB struct {
	client *dynamodb.DynamoDB
}

// DynamoClient create dynamo client.
func (d *DynamoDB) DynamoClient() (DynamoAPI, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	d.client = dynamodb.New(sess)

	return d.client, nil
}
