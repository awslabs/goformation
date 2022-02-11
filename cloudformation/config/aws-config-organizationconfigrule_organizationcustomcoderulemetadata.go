package config

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// OrganizationConfigRule_OrganizationCustomCodeRuleMetadata AWS CloudFormation Resource (AWS::Config::OrganizationConfigRule.OrganizationCustomCodeRuleMetadata)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html
type OrganizationConfigRule_OrganizationCustomCodeRuleMetadata struct {

	// CodeText AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-codetext
	CodeText string `json:"CodeText"`

	// DebugLogDeliveryAccounts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-debuglogdeliveryaccounts
	DebugLogDeliveryAccounts *[]string `json:"DebugLogDeliveryAccounts,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-description
	Description *string `json:"Description,omitempty"`

	// InputParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-inputparameters
	InputParameters *string `json:"InputParameters,omitempty"`

	// MaximumExecutionFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-maximumexecutionfrequency
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty"`

	// OrganizationConfigRuleTriggerTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-organizationconfigruletriggertypes
	OrganizationConfigRuleTriggerTypes *[]string `json:"OrganizationConfigRuleTriggerTypes,omitempty"`

	// ResourceIdScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-resourceidscope
	ResourceIdScope *string `json:"ResourceIdScope,omitempty"`

	// ResourceTypesScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-resourcetypesscope
	ResourceTypesScope *[]string `json:"ResourceTypesScope,omitempty"`

	// Runtime AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-runtime
	Runtime string `json:"Runtime"`

	// TagKeyScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-tagkeyscope
	TagKeyScope *string `json:"TagKeyScope,omitempty"`

	// TagValueScope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustomcoderulemetadata.html#cfn-config-organizationconfigrule-organizationcustomcoderulemetadata-tagvaluescope
	TagValueScope *string `json:"TagValueScope,omitempty"`

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
func (r *OrganizationConfigRule_OrganizationCustomCodeRuleMetadata) AWSCloudFormationType() string {
	return "AWS::Config::OrganizationConfigRule.OrganizationCustomCodeRuleMetadata"
}
