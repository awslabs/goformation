package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.Projection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
type AWSDynamoDBTable_Projection struct {

	// NonKeyAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-nonkeyatt
	NonKeyAttributes []string `json:"NonKeyAttributes"`

	// ProjectionType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html#cfn-dynamodb-projectionobj-projtype
	ProjectionType string `json:"ProjectionType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_Projection) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.Projection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_Projection) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDynamoDBTable_ProjectionResources retrieves all AWSDynamoDBTable_Projection items from a CloudFormation template
func GetAllAWSDynamoDBTable_Projection(template *Template) map[string]*AWSDynamoDBTable_Projection {

	results := map[string]*AWSDynamoDBTable_Projection{}
	for name, resource := range template.Resources {
		result := &AWSDynamoDBTable_Projection{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDynamoDBTable_ProjectionWithName retrieves all AWSDynamoDBTable_Projection items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDynamoDBTable_Projection(name string, template *Template) (*AWSDynamoDBTable_Projection, error) {

	result := &AWSDynamoDBTable_Projection{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDynamoDBTable_Projection{}, errors.New("resource not found")

}
