// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_TimeRangeDrillDownFilter AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TimeRangeDrillDownFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-timerangedrilldownfilter.html
type Dashboard_TimeRangeDrillDownFilter struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-timerangedrilldownfilter.html#cfn-quicksight-dashboard-timerangedrilldownfilter-column
	Column *Dashboard_ColumnIdentifier `json:"Column"`

	// RangeMaximum AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-timerangedrilldownfilter.html#cfn-quicksight-dashboard-timerangedrilldownfilter-rangemaximum
	RangeMaximum string `json:"RangeMaximum"`

	// RangeMinimum AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-timerangedrilldownfilter.html#cfn-quicksight-dashboard-timerangedrilldownfilter-rangeminimum
	RangeMinimum string `json:"RangeMinimum"`

	// TimeGranularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-timerangedrilldownfilter.html#cfn-quicksight-dashboard-timerangedrilldownfilter-timegranularity
	TimeGranularity string `json:"TimeGranularity"`

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
func (r *Dashboard_TimeRangeDrillDownFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TimeRangeDrillDownFilter"
}
