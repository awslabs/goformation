package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DynamoDB::Table.GlobalSecondaryIndex AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type AWSDynamoDBTable_GlobalSecondaryIndex struct {

	// IndexName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-indexname
	IndexName string `json:"IndexName"`

	// KeySchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-keyschema
	KeySchema []AWSDynamoDBTable_KeySchema `json:"KeySchema"`

	// Projection AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-projection
	Projection AWSDynamoDBTable_Projection `json:"Projection"`

	// ProvisionedThroughput AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-provisionedthroughput
	ProvisionedThroughput AWSDynamoDBTable_ProvisionedThroughput `json:"ProvisionedThroughput"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_GlobalSecondaryIndex) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.GlobalSecondaryIndex"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDynamoDBTable_GlobalSecondaryIndex) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
