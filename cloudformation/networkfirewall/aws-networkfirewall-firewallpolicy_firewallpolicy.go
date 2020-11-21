package networkfirewall

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// FirewallPolicy_FirewallPolicy AWS CloudFormation Resource (AWS::NetworkFirewall::FirewallPolicy.FirewallPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html
type FirewallPolicy_FirewallPolicy struct {

	// StatefulRuleGroupReferences AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy-statefulrulegroupreferences
	StatefulRuleGroupReferences *FirewallPolicy_StatefulRuleGroupReferences `json:"StatefulRuleGroupReferences,omitempty"`

	// StatelessCustomActions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy-statelesscustomactions
	StatelessCustomActions *FirewallPolicy_CustomActions `json:"StatelessCustomActions,omitempty"`

	// StatelessDefaultActions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy-statelessdefaultactions
	StatelessDefaultActions *FirewallPolicy_StatelessActions `json:"StatelessDefaultActions,omitempty"`

	// StatelessFragmentDefaultActions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy-statelessfragmentdefaultactions
	StatelessFragmentDefaultActions *FirewallPolicy_StatelessActions `json:"StatelessFragmentDefaultActions,omitempty"`

	// StatelessRuleGroupReferences AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html#cfn-networkfirewall-firewallpolicy-firewallpolicy-statelessrulegroupreferences
	StatelessRuleGroupReferences *FirewallPolicy_StatelessRuleGroupReferences `json:"StatelessRuleGroupReferences,omitempty"`

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
func (r *FirewallPolicy_FirewallPolicy) AWSCloudFormationType() string {
	return "AWS::NetworkFirewall::FirewallPolicy.FirewallPolicy"
}
