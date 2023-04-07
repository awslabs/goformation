// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_TableOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TableOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableoptions.html
type Dashboard_TableOptions struct {

	// CellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableoptions.html#cfn-quicksight-dashboard-tableoptions-cellstyle
	CellStyle *Dashboard_TableCellStyle `json:"CellStyle,omitempty"`

	// HeaderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableoptions.html#cfn-quicksight-dashboard-tableoptions-headerstyle
	HeaderStyle *Dashboard_TableCellStyle `json:"HeaderStyle,omitempty"`

	// Orientation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableoptions.html#cfn-quicksight-dashboard-tableoptions-orientation
	Orientation *string `json:"Orientation,omitempty"`

	// RowAlternateColorOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableoptions.html#cfn-quicksight-dashboard-tableoptions-rowalternatecoloroptions
	RowAlternateColorOptions *Dashboard_RowAlternateColorOptions `json:"RowAlternateColorOptions,omitempty"`

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
func (r *Dashboard_TableOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TableOptions"
}