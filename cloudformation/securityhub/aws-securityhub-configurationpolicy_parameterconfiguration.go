// Code generated by "go generate". Please don't change this file directly.

package securityhub

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ConfigurationPolicy_ParameterConfiguration AWS CloudFormation Resource (AWS::SecurityHub::ConfigurationPolicy.ParameterConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-parameterconfiguration.html
type ConfigurationPolicy_ParameterConfiguration struct {

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-parameterconfiguration.html#cfn-securityhub-configurationpolicy-parameterconfiguration-value
	Value *ConfigurationPolicy_ParameterValue `json:"Value,omitempty"`

	// ValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-parameterconfiguration.html#cfn-securityhub-configurationpolicy-parameterconfiguration-valuetype
	ValueType string `json:"ValueType"`

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
func (r *ConfigurationPolicy_ParameterConfiguration) AWSCloudFormationType() string {
	return "AWS::SecurityHub::ConfigurationPolicy.ParameterConfiguration"
}