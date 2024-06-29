// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_FilterListControl AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FilterListControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html
type Dashboard_FilterListControl struct {

	// CascadingControlConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-cascadingcontrolconfiguration
	CascadingControlConfiguration *Dashboard_CascadingControlConfiguration `json:"CascadingControlConfiguration,omitempty"`

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-displayoptions
	DisplayOptions *Dashboard_ListControlDisplayOptions `json:"DisplayOptions,omitempty"`

	// FilterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-filtercontrolid
	FilterControlId string `json:"FilterControlId"`

	// SelectableValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-selectablevalues
	SelectableValues *Dashboard_FilterSelectableValues `json:"SelectableValues,omitempty"`

	// SourceFilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-sourcefilterid
	SourceFilterId string `json:"SourceFilterId"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-title
	Title string `json:"Title"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterlistcontrol.html#cfn-quicksight-dashboard-filterlistcontrol-type
	Type *string `json:"Type,omitempty"`

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
func (r *Dashboard_FilterListControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FilterListControl"
}