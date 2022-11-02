package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"mutant-project/models"
)

// table with the human data.
const tableName string = "humans_dna"

// MutantRepository struct for mutant repository methods.
type MutantRepository struct {
	Client *dynamodb.DynamoDB
	Table  string
}

// MutantRepositoryInterface interface for repositories.MutantRepository
// Save human dna info and if this contains dna mutant or not.
// GetDNAStats obtain the stats of the human validate.
type MutantRepositoryInterface interface {
	Save(dna []string, isMutant bool) error
	GetDNAStats() (totalMutants int, totalHumans int, err error)
}

// Save human dna info and if this contains dna mutant or not.
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

// GetDNAStats obtain the stats of the human validate.
// Count total humans validate.
// Count total humans with dna mutant.
func (r *MutantRepository) GetDNAStats() (totalMutants int, totalHumans int, err error) {
	// obtain human data.
	result, err := r.Client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(r.Table),
	})

	if err != nil {
		return
	}

	// loop data and count total humans and total humans with dna mutant.
	for _, data := range result.Items {
		totalHumans++

		// obtain data from dynamo result.
		item := models.HumanDNA{}
		err = dynamodbattribute.UnmarshalMap(data, &item)
		if err != nil {
			return
		}

		// check if the human has dna mutant.
		if item.IsMutant {
			totalMutants++
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
		Table:  tableName,
	}
}
