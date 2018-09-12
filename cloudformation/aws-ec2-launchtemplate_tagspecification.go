package cloudformation

import (
	"encoding/json"
)

// AWSEC2LaunchTemplate_TagSpecification AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.TagSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html
type AWSEC2LaunchTemplate_TagSpecification struct {

	// ResourceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html#cfn-ec2-launchtemplate-tagspecification-resourcetype
	ResourceType *Value `json:"ResourceType,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html#cfn-ec2-launchtemplate-tagspecification-tags
	Tags []Tag `json:"Tags,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_TagSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.TagSpecification"
}

func (r *AWSEC2LaunchTemplate_TagSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
