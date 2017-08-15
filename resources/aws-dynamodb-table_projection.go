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
