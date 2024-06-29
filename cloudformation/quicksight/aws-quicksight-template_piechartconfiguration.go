// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_PieChartConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.PieChartConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html
type Template_PieChartConfiguration struct {

	// CategoryLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-categorylabeloptions
	CategoryLabelOptions *Template_ChartAxisLabelOptions `json:"CategoryLabelOptions,omitempty"`

	// ContributionAnalysisDefaults AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-contributionanalysisdefaults
	ContributionAnalysisDefaults []Template_ContributionAnalysisDefault `json:"ContributionAnalysisDefaults,omitempty"`

	// DataLabels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-datalabels
	DataLabels *Template_DataLabelOptions `json:"DataLabels,omitempty"`

	// DonutOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-donutoptions
	DonutOptions *Template_DonutOptions `json:"DonutOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-fieldwells
	FieldWells *Template_PieChartFieldWells `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-legend
	Legend *Template_LegendOptions `json:"Legend,omitempty"`

	// SmallMultiplesOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-smallmultiplesoptions
	SmallMultiplesOptions *Template_SmallMultiplesOptions `json:"SmallMultiplesOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-sortconfiguration
	SortConfiguration *Template_PieChartSortConfiguration `json:"SortConfiguration,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-tooltip
	Tooltip *Template_TooltipOptions `json:"Tooltip,omitempty"`

	// ValueLabelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-valuelabeloptions
	ValueLabelOptions *Template_ChartAxisLabelOptions `json:"ValueLabelOptions,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-piechartconfiguration.html#cfn-quicksight-template-piechartconfiguration-visualpalette
	VisualPalette *Template_VisualPalette `json:"VisualPalette,omitempty"`

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
func (r *Template_PieChartConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PieChartConfiguration"
}