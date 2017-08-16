package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSDynamoDBTable AWS CloudFormation Resource (AWS::DynamoDB::Table)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
type AWSDynamoDBTable struct {

	// AttributeDefinitions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-attributedef
	AttributeDefinitions []AWSDynamoDBTable_AttributeDefinition `json:"AttributeDefinitions"`

	// GlobalSecondaryIndexes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-gsi
	GlobalSecondaryIndexes []AWSDynamoDBTable_GlobalSecondaryIndex `json:"GlobalSecondaryIndexes"`

	// KeySchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-keyschema
	KeySchema []AWSDynamoDBTable_KeySchema `json:"KeySchema"`

	// LocalSecondaryIndexes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-lsi
	LocalSecondaryIndexes []AWSDynamoDBTable_LocalSecondaryIndex `json:"LocalSecondaryIndexes"`

	// ProvisionedThroughput AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-provisionedthroughput
	ProvisionedThroughput AWSDynamoDBTable_ProvisionedThroughput `json:"ProvisionedThroughput"`

	// StreamSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-streamspecification
	StreamSpecification AWSDynamoDBTable_StreamSpecification `json:"StreamSpecification"`

	// TableName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tablename
	TableName string `json:"TableName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTableResources retrieves all AWSDynamoDBTable items from a CloudFormation template
func (t *Template) GetAllAWSDynamoDBTableResources() map[string]*AWSDynamoDBTable {

	results := map[string]*AWSDynamoDBTable{}
	for name, resource := range t.Resources {
		result := &AWSDynamoDBTable{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTableWithName retrieves all AWSDynamoDBTable items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDynamoDBTableWithName(name string) (*AWSDynamoDBTable, error) {

	result := &AWSDynamoDBTable{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable{}, errors.New("resource not found")

}
