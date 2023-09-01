// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_PivotTableOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.PivotTableOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html
type Analysis_PivotTableOptions struct {

	// CellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-cellstyle
	CellStyle *Analysis_TableCellStyle `json:"CellStyle,omitempty"`

	// CollapsedRowDimensionsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-collapsedrowdimensionsvisibility
	CollapsedRowDimensionsVisibility *string `json:"CollapsedRowDimensionsVisibility,omitempty"`

	// ColumnHeaderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-columnheaderstyle
	ColumnHeaderStyle *Analysis_TableCellStyle `json:"ColumnHeaderStyle,omitempty"`

	// ColumnNamesVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-columnnamesvisibility
	ColumnNamesVisibility *string `json:"ColumnNamesVisibility,omitempty"`

	// DefaultCellWidth AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-defaultcellwidth
	DefaultCellWidth *string `json:"DefaultCellWidth,omitempty"`

	// MetricPlacement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-metricplacement
	MetricPlacement *string `json:"MetricPlacement,omitempty"`

	// RowAlternateColorOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-rowalternatecoloroptions
	RowAlternateColorOptions *Analysis_RowAlternateColorOptions `json:"RowAlternateColorOptions,omitempty"`

	// RowFieldNamesStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-rowfieldnamesstyle
	RowFieldNamesStyle *Analysis_TableCellStyle `json:"RowFieldNamesStyle,omitempty"`

	// RowHeaderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-rowheaderstyle
	RowHeaderStyle *Analysis_TableCellStyle `json:"RowHeaderStyle,omitempty"`

	// RowsLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-rowslabeloptions
	RowsLabelOptions *Analysis_PivotTableRowsLabelOptions `json:"RowsLabelOptions,omitempty"`

	// RowsLayout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-rowslayout
	RowsLayout *string `json:"RowsLayout,omitempty"`

	// SingleMetricVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-singlemetricvisibility
	SingleMetricVisibility *string `json:"SingleMetricVisibility,omitempty"`

	// ToggleButtonsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottableoptions.html#cfn-quicksight-analysis-pivottableoptions-togglebuttonsvisibility
	ToggleButtonsVisibility *string `json:"ToggleButtonsVisibility,omitempty"`

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
func (r *Analysis_PivotTableOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PivotTableOptions"
}
