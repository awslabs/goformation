package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.AttributeDefinition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
type AWSDynamoDBTable_AttributeDefinition struct {

	// AttributeName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename
	AttributeName string `json:"AttributeName"`

	// AttributeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html#cfn-dynamodb-attributedef-attributename-attributetype
	AttributeType string `json:"AttributeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_AttributeDefinition) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.AttributeDefinition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_AttributeDefinition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTable_AttributeDefinitionResources retrieves all AWSDynamoDBTable_AttributeDefinition items from a CloudFormation template
func GetAllAWSDynamoDBTable_AttributeDefinition(template *Template) map[string]*AWSDynamoDBTable_AttributeDefinition {

	results := map[string]*AWSDynamoDBTable_AttributeDefinition{}
	for name, resource := range template.Resources {
		result := &AWSDynamoDBTable_AttributeDefinition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTable_AttributeDefinitionWithName retrieves all AWSDynamoDBTable_AttributeDefinition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDynamoDBTable_AttributeDefinition(name string, template *Template) (*AWSDynamoDBTable_AttributeDefinition, error) {

	result := &AWSDynamoDBTable_AttributeDefinition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable_AttributeDefinition{}, errors.New("resource not found")

}
