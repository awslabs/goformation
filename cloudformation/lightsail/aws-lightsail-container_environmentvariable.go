// Code generated by "go generate". Please don't change this file directly.

package lightsail

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Container_EnvironmentVariable AWS CloudFormation Resource (AWS::Lightsail::Container.EnvironmentVariable)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-environmentvariable.html
type Container_EnvironmentVariable struct {

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-environmentvariable.html#cfn-lightsail-container-environmentvariable-value
	Value *string `json:"Value,omitempty"`

	// Variable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-environmentvariable.html#cfn-lightsail-container-environmentvariable-variable
	Variable *string `json:"Variable,omitempty"`

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
func (r *Container_EnvironmentVariable) AWSCloudFormationType() string {
	return "AWS::Lightsail::Container.EnvironmentVariable"
}
