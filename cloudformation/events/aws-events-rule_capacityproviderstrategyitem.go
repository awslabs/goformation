package events

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Rule_CapacityProviderStrategyItem AWS CloudFormation Resource (AWS::Events::Rule.CapacityProviderStrategyItem)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-capacityproviderstrategyitem.html
type Rule_CapacityProviderStrategyItem struct {

	// Base AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-capacityproviderstrategyitem.html#cfn-events-rule-capacityproviderstrategyitem-base
	Base *int `json:"Base,omitempty"`

	// CapacityProvider AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-capacityproviderstrategyitem.html#cfn-events-rule-capacityproviderstrategyitem-capacityprovider
	CapacityProvider string `json:"CapacityProvider"`

	// Weight AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-capacityproviderstrategyitem.html#cfn-events-rule-capacityproviderstrategyitem-weight
	Weight *int `json:"Weight,omitempty"`

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
func (r *Rule_CapacityProviderStrategyItem) AWSCloudFormationType() string {
	return "AWS::Events::Rule.CapacityProviderStrategyItem"
}
