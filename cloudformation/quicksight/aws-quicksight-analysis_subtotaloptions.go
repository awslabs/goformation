// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_SubtotalOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.SubtotalOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html
type Analysis_SubtotalOptions struct {

	// CustomLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-customlabel
	CustomLabel *string `json:"CustomLabel,omitempty"`

	// FieldLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-fieldlevel
	FieldLevel *string `json:"FieldLevel,omitempty"`

	// FieldLevelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-fieldleveloptions
	FieldLevelOptions []Analysis_PivotTableFieldSubtotalOptions `json:"FieldLevelOptions,omitempty"`

	// MetricHeaderCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-metricheadercellstyle
	MetricHeaderCellStyle *Analysis_TableCellStyle `json:"MetricHeaderCellStyle,omitempty"`

	// StyleTargets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-styletargets
	StyleTargets []Analysis_TableStyleTarget `json:"StyleTargets,omitempty"`

	// TotalCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-totalcellstyle
	TotalCellStyle *Analysis_TableCellStyle `json:"TotalCellStyle,omitempty"`

	// TotalsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-totalsvisibility
	TotalsVisibility *string `json:"TotalsVisibility,omitempty"`

	// ValueCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-valuecellstyle
	ValueCellStyle *Analysis_TableCellStyle `json:"ValueCellStyle,omitempty"`

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
func (r *Analysis_SubtotalOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.SubtotalOptions"
}