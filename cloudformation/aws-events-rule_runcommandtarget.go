package cloudformation

import (
	"encoding/json"
)

// AWSEventsRule_RunCommandTarget AWS CloudFormation Resource (AWS::Events::Rule.RunCommandTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html
type AWSEventsRule_RunCommandTarget struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html#cfn-events-rule-runcommandtarget-key
	Key *Value `json:"Key,omitempty"`

	// Values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html#cfn-events-rule-runcommandtarget-values
	Values []*Value `json:"Values,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_RunCommandTarget) AWSCloudFormationType() string {
	return "AWS::Events::Rule.RunCommandTarget"
}

func (r *AWSEventsRule_RunCommandTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
