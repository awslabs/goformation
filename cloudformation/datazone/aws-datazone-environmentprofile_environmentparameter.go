// Code generated by "go generate". Please don't change this file directly.

package datazone

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// EnvironmentProfile_EnvironmentParameter AWS CloudFormation Resource (AWS::DataZone::EnvironmentProfile.EnvironmentParameter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentprofile-environmentparameter.html
type EnvironmentProfile_EnvironmentParameter struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentprofile-environmentparameter.html#cfn-datazone-environmentprofile-environmentparameter-name
	Name *string `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-environmentprofile-environmentparameter.html#cfn-datazone-environmentprofile-environmentparameter-value
	Value *string `json:"Value,omitempty"`

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
func (r *EnvironmentProfile_EnvironmentParameter) AWSCloudFormationType() string {
	return "AWS::DataZone::EnvironmentProfile.EnvironmentParameter"
}
