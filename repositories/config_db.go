package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Dynamo struct{}

// DynamoAPI interface for DynamoDB methods.
type DynamoAPI interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
}

func (d *Dynamo) DynamoClient() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("com.amazonaws.us-east-1.dynamodb\t"),
	}

	sess := session.Must(session.NewSession(config))

	return dynamodb.New(sess)
}

// DynamoProvider interface for Dynamo client.
type DynamoProvider interface {
	DynamoClient() (DynamoAPI, error)
}

type Config struct {
	EndPoint string
	Region   string
}

type DynamoDB struct {
	client *dynamodb.DynamoDB
	config Config
}

func (d *DynamoDB) DynamoClient() (DynamoAPI, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Endpoint: &d.config.EndPoint,
		Region:   &d.config.Region,
	})

	if err != nil {
		return nil, err
	}

	d.client = dynamodb.New(sess)
	return d.client, nil
}
