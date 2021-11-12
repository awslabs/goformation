package ec2

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// EC2Fleet_BaselineEbsBandwidthMbpsRequest AWS CloudFormation Resource (AWS::EC2::EC2Fleet.BaselineEbsBandwidthMbpsRequest)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-baselineebsbandwidthmbpsrequest.html
type EC2Fleet_BaselineEbsBandwidthMbpsRequest struct {

	// Max AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-baselineebsbandwidthmbpsrequest.html#cfn-ec2-ec2fleet-baselineebsbandwidthmbpsrequest-max
	Max int `json:"Max,omitempty"`

	// Min AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-baselineebsbandwidthmbpsrequest.html#cfn-ec2-ec2fleet-baselineebsbandwidthmbpsrequest-min
	Min int `json:"Min,omitempty"`

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
func (r *EC2Fleet_BaselineEbsBandwidthMbpsRequest) AWSCloudFormationType() string {
	return "AWS::EC2::EC2Fleet.BaselineEbsBandwidthMbpsRequest"
}
