package cloudformation

import (
	"encoding/json"
)

// AWSDynamoDBTable_StreamSpecification AWS CloudFormation Resource (AWS::DynamoDB::Table.StreamSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html
type AWSDynamoDBTable_StreamSpecification struct {

	// StreamViewType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html#cfn-dynamodb-streamspecification-streamviewtype
	StreamViewType *Value `json:"StreamViewType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_StreamSpecification) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.StreamSpecification"
}

func (r *AWSDynamoDBTable_StreamSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
