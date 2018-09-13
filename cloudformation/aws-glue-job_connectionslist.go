package cloudformation

import (
	"encoding/json"
)

// AWSGlueJob_ConnectionsList AWS CloudFormation Resource (AWS::Glue::Job.ConnectionsList)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-connectionslist.html
type AWSGlueJob_ConnectionsList struct {

	// Connections AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-connectionslist.html#cfn-glue-job-connectionslist-connections
	Connections []*Value `json:"Connections,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueJob_ConnectionsList) AWSCloudFormationType() string {
	return "AWS::Glue::Job.ConnectionsList"
}

func (r *AWSGlueJob_ConnectionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
