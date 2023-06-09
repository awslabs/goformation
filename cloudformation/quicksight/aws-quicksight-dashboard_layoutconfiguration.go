// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_LayoutConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.LayoutConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html
type Dashboard_LayoutConfiguration struct {

	// FreeFormLayout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html#cfn-quicksight-dashboard-layoutconfiguration-freeformlayout
	FreeFormLayout *Dashboard_FreeFormLayoutConfiguration `json:"FreeFormLayout,omitempty"`

	// GridLayout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html#cfn-quicksight-dashboard-layoutconfiguration-gridlayout
	GridLayout *Dashboard_GridLayoutConfiguration `json:"GridLayout,omitempty"`

	// SectionBasedLayout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html#cfn-quicksight-dashboard-layoutconfiguration-sectionbasedlayout
	SectionBasedLayout *Dashboard_SectionBasedLayoutConfiguration `json:"SectionBasedLayout,omitempty"`

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
func (r *Dashboard_LayoutConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.LayoutConfiguration"
}