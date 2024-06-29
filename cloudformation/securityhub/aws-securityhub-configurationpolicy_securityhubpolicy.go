// Code generated by "go generate". Please don't change this file directly.

package securityhub

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ConfigurationPolicy_SecurityHubPolicy AWS CloudFormation Resource (AWS::SecurityHub::ConfigurationPolicy.SecurityHubPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html
type ConfigurationPolicy_SecurityHubPolicy struct {

	// EnabledStandardIdentifiers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html#cfn-securityhub-configurationpolicy-securityhubpolicy-enabledstandardidentifiers
	EnabledStandardIdentifiers []string `json:"EnabledStandardIdentifiers,omitempty"`

	// SecurityControlsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html#cfn-securityhub-configurationpolicy-securityhubpolicy-securitycontrolsconfiguration
	SecurityControlsConfiguration *ConfigurationPolicy_SecurityControlsConfiguration `json:"SecurityControlsConfiguration,omitempty"`

	// ServiceEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html#cfn-securityhub-configurationpolicy-securityhubpolicy-serviceenabled
	ServiceEnabled *bool `json:"ServiceEnabled,omitempty"`

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
func (r *ConfigurationPolicy_SecurityHubPolicy) AWSCloudFormationType() string {
	return "AWS::SecurityHub::ConfigurationPolicy.SecurityHubPolicy"
}