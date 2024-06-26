// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_PivotTableConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.PivotTableConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html
type Dashboard_PivotTableConfiguration struct {

	// FieldOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-fieldoptions
	FieldOptions *Dashboard_PivotTableFieldOptions `json:"FieldOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-fieldwells
	FieldWells *Dashboard_PivotTableFieldWells `json:"FieldWells,omitempty"`

	// PaginatedReportOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-paginatedreportoptions
	PaginatedReportOptions *Dashboard_PivotTablePaginatedReportOptions `json:"PaginatedReportOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-sortconfiguration
	SortConfiguration *Dashboard_PivotTableSortConfiguration `json:"SortConfiguration,omitempty"`

	// TableOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-tableoptions
	TableOptions *Dashboard_PivotTableOptions `json:"TableOptions,omitempty"`

	// TotalOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-totaloptions
	TotalOptions *Dashboard_PivotTableTotalOptions `json:"TotalOptions,omitempty"`

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
func (r *Dashboard_PivotTableConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.PivotTableConfiguration"
}
