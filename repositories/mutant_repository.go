package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// MutantRepository ...
type MutantRepository struct {
	Client *dynamodb.DynamoDB
	Table  string
}

// Save ...
func (r *MutantRepository) Save(dna []string) error {
	_, err := r.Client.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("pk"),
			},
			"sk": {
				S: aws.String("sk"),
			},
		},
		TableName: aws.String(r.Table),
	})

	if err != nil {
		return err
	}

	return nil
}

func NewMutantRepository() *MutantRepository {
	dynamoProvider := DynamoDB{}
	clientDynamo, err := dynamoProvider.DynamoClient()
	if err != nil {
		panic("dynamo client error")
	}

	return &MutantRepository{
		Client: clientDynamo.(*dynamodb.DynamoDB),
		Table:  "mutants",
	}
}
