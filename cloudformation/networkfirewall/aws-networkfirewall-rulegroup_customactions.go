package networkfirewall

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// RuleGroup_CustomActions AWS CloudFormation Resource (AWS::NetworkFirewall::RuleGroup.CustomActions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customactions.html
type RuleGroup_CustomActions struct {

	// CustomActions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-customactions.html#cfn-networkfirewall-rulegroup-customactions-customactions
	CustomActions []RuleGroup_CustomAction `json:"CustomActions,omitempty"`

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
func (r *RuleGroup_CustomActions) AWSCloudFormationType() string {
	return "AWS::NetworkFirewall::RuleGroup.CustomActions"
}
