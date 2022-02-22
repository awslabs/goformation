// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// LaunchTemplate_BaselineEbsBandwidthMbps AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.BaselineEbsBandwidthMbps)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineebsbandwidthmbps.html
type LaunchTemplate_BaselineEbsBandwidthMbps struct {

	// Max AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineebsbandwidthmbps.html#cfn-ec2-launchtemplate-baselineebsbandwidthmbps-max
	Max *int `json:"Max,omitempty"`

	// Min AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-baselineebsbandwidthmbps.html#cfn-ec2-launchtemplate-baselineebsbandwidthmbps-min
	Min *int `json:"Min,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *LaunchTemplate_BaselineEbsBandwidthMbps) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.BaselineEbsBandwidthMbps"
}
