// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_FilterListConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FilterListConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistconfiguration.html
type Dashboard_FilterListConfiguration struct {

	// CategoryValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistconfiguration.html#cfn-quicksight-dashboard-filterlistconfiguration-categoryvalues
	CategoryValues []string `json:"CategoryValues,omitempty"`

	// MatchOperator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistconfiguration.html#cfn-quicksight-dashboard-filterlistconfiguration-matchoperator
	MatchOperator string `json:"MatchOperator"`

	// NullOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistconfiguration.html#cfn-quicksight-dashboard-filterlistconfiguration-nulloption
	NullOption *string `json:"NullOption,omitempty"`

	// SelectAllOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistconfiguration.html#cfn-quicksight-dashboard-filterlistconfiguration-selectalloptions
	SelectAllOptions *string `json:"SelectAllOptions,omitempty"`

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
func (r *Dashboard_FilterListConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FilterListConfiguration"
}
