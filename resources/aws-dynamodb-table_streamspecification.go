package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.StreamSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html
type AWSDynamoDBTable_StreamSpecification struct {

	// StreamViewType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html#cfn-dynamodb-streamspecification-streamviewtype

	StreamViewType string `json:"StreamViewType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_StreamSpecification) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.StreamSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_StreamSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTable_StreamSpecificationResources retrieves all AWSDynamoDBTable_StreamSpecification items from a CloudFormation template
func GetAllAWSDynamoDBTable_StreamSpecification(template *Template) map[string]*AWSDynamoDBTable_StreamSpecification {

	results := map[string]*AWSDynamoDBTable_StreamSpecification{}
	for name, resource := range template.Resources {
		result := &AWSDynamoDBTable_StreamSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTable_StreamSpecificationWithName retrieves all AWSDynamoDBTable_StreamSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDynamoDBTable_StreamSpecification(name string, template *Template) (*AWSDynamoDBTable_StreamSpecification, error) {

	result := &AWSDynamoDBTable_StreamSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable_StreamSpecification{}, errors.New("resource not found")

}
