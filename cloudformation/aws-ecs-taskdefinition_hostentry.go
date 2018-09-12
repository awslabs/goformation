package cloudformation

import (
	"encoding/json"
)

// AWSECSTaskDefinition_HostEntry AWS CloudFormation Resource (AWS::ECS::TaskDefinition.HostEntry)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html
type AWSECSTaskDefinition_HostEntry struct {

	// Hostname AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-hostname
	Hostname *Value `json:"Hostname,omitempty"`

	// IpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-ipaddress
	IpAddress *Value `json:"IpAddress,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_HostEntry) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.HostEntry"
}

func (r *AWSECSTaskDefinition_HostEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
