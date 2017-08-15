package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.KeySchema AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
type AWSDynamoDBTable_KeySchema struct {

	// AttributeName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-attributename
	AttributeName string `json:"AttributeName"`

	// KeyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html#aws-properties-dynamodb-keyschema-keytype
	KeyType string `json:"KeyType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_KeySchema) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.KeySchema"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_KeySchema) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTable_KeySchemaResources retrieves all AWSDynamoDBTable_KeySchema items from a CloudFormation template
func GetAllAWSDynamoDBTable_KeySchema(template *Template) map[string]*AWSDynamoDBTable_KeySchema {

	results := map[string]*AWSDynamoDBTable_KeySchema{}
	for name, resource := range template.Resources {
		result := &AWSDynamoDBTable_KeySchema{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTable_KeySchemaWithName retrieves all AWSDynamoDBTable_KeySchema items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDynamoDBTable_KeySchema(name string, template *Template) (*AWSDynamoDBTable_KeySchema, error) {

	result := &AWSDynamoDBTable_KeySchema{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable_KeySchema{}, errors.New("resource not found")

}
