package networkfirewall

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// RuleGroup_PortRanges AWS CloudFormation Resource (AWS::NetworkFirewall::RuleGroup.PortRanges)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portranges.html
type RuleGroup_PortRanges struct {

	// PortRanges AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portranges.html#cfn-networkfirewall-rulegroup-portranges-portranges
	PortRanges []RuleGroup_PortRange `json:"PortRanges,omitempty"`

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
func (r *RuleGroup_PortRanges) AWSCloudFormationType() string {
	return "AWS::NetworkFirewall::RuleGroup.PortRanges"
}
