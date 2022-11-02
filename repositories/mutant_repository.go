package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"mutant-project/models"
)

// MutantRepository struct for mutant repository methods.
type MutantRepository struct {
	Client *dynamodb.DynamoDB
	Table  string
}

type MutantRepositoryInterface interface {
	Save(dna []string, isMutant bool) error
	GetDNAStats() (totalMutants int, totalHumans int, err error)
}

// Save dna human info.
func (r *MutantRepository) Save(dna []string, isMutant bool) error {
	// generate partition key with human dna.
	pk := ""
	for i := 0; i < len(dna); i++ {
		pk = pk + dna[i] + "-"
	}

	// save data into dynamo db table
	_, err := r.Client.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(pk[:len(pk)-1]),
			},
			"is_mutant": {
				BOOL: aws.Bool(isMutant),
			},
		},
		TableName: aws.String(r.Table),
	})

	// check error.
	if err != nil {
		return err
	}

	return nil
}

// GetDNAStats ...
func (r *MutantRepository) GetDNAStats() (totalMutants int, totalHumans int, err error) {
	result, err := r.Client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(r.Table),
	})

	if err != nil {
		return
	}

	for _, data := range result.Items {
		item := models.Mutant{}
		err = dynamodbattribute.UnmarshalMap(data, &item)
		if err != nil {
			return
		}

		if item.IsMutant {
			totalMutants++
		} else {
			totalHumans++
		}
	}

	return
}

// NewMutantRepository create repository arguments.
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
