package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.LocalSecondaryIndex AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html
type AWSDynamoDBTable_LocalSecondaryIndex struct {

	// IndexName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-indexname
	IndexName string `json:"IndexName"`

	// KeySchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-keyschema
	KeySchema []AWSDynamoDBTable_KeySchema `json:"KeySchema"`

	// Projection AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html#cfn-dynamodb-lsi-projection
	Projection AWSDynamoDBTable_Projection `json:"Projection"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_LocalSecondaryIndex) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.LocalSecondaryIndex"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_LocalSecondaryIndex) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTable_LocalSecondaryIndexResources retrieves all AWSDynamoDBTable_LocalSecondaryIndex items from a CloudFormation template
func GetAllAWSDynamoDBTable_LocalSecondaryIndex(template *Template) map[string]*AWSDynamoDBTable_LocalSecondaryIndex {

	results := map[string]*AWSDynamoDBTable_LocalSecondaryIndex{}
	for name, resource := range template.Resources {
		result := &AWSDynamoDBTable_LocalSecondaryIndex{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTable_LocalSecondaryIndexWithName retrieves all AWSDynamoDBTable_LocalSecondaryIndex items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDynamoDBTable_LocalSecondaryIndex(name string, template *Template) (*AWSDynamoDBTable_LocalSecondaryIndex, error) {

	result := &AWSDynamoDBTable_LocalSecondaryIndex{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable_LocalSecondaryIndex{}, errors.New("resource not found")

}
