package cloudformation

import (
	"encoding/json"
)

// AWSServerlessFunction_SNSEvent AWS CloudFormation Resource (AWS::Serverless::Function.SNSEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sns
type AWSServerlessFunction_SNSEvent struct {

	// Topic AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sns
	Topic *Value `json:"Topic,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_SNSEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.SNSEvent"
}

func (r *AWSServerlessFunction_SNSEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
