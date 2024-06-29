// Code generated by "go generate". Please don't change this file directly.

package fms

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Policy_NetworkAclCommonPolicy AWS CloudFormation Resource (AWS::FMS::Policy.NetworkAclCommonPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclcommonpolicy.html
type Policy_NetworkAclCommonPolicy struct {

	// NetworkAclEntrySet AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclcommonpolicy.html#cfn-fms-policy-networkaclcommonpolicy-networkaclentryset
	NetworkAclEntrySet *Policy_NetworkAclEntrySet `json:"NetworkAclEntrySet"`

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
func (r *Policy_NetworkAclCommonPolicy) AWSCloudFormationType() string {
	return "AWS::FMS::Policy.NetworkAclCommonPolicy"
}