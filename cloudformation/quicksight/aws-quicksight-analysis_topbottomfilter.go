// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_TopBottomFilter AWS CloudFormation Resource (AWS::QuickSight::Analysis.TopBottomFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html
type Analysis_TopBottomFilter struct {

	// AggregationSortConfigurations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-aggregationsortconfigurations
	AggregationSortConfigurations []Analysis_AggregationSortConfiguration `json:"AggregationSortConfigurations"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-column
	Column *Analysis_ColumnIdentifier `json:"Column"`

	// DefaultFilterControlConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-defaultfiltercontrolconfiguration
	DefaultFilterControlConfiguration *Analysis_DefaultFilterControlConfiguration `json:"DefaultFilterControlConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-filterid
	FilterId string `json:"FilterId"`

	// Limit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-limit
	Limit *float64 `json:"Limit,omitempty"`

	// ParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-parametername
	ParameterName *string `json:"ParameterName,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottomfilter.html#cfn-quicksight-analysis-topbottomfilter-timegranularity
	TimeGranularity *string `json:"TimeGranularity,omitempty"`

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
func (r *Analysis_TopBottomFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TopBottomFilter"
}