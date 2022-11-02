package repositories

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// MutantRepository ...
type MutantRepository struct {
	Client *dynamodb.DynamoDB
	Table  string
}

// Save ...
func (s *MutantRepository) Save(dna []string) error {
	return nil
}
