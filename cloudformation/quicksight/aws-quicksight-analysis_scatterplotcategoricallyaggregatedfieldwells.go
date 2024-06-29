// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_ScatterPlotCategoricallyAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Analysis.ScatterPlotCategoricallyAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells.html
type Analysis_ScatterPlotCategoricallyAggregatedFieldWells struct {

	// Category AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells-category
	Category []Analysis_DimensionField `json:"Category,omitempty"`

	// Label AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells-label
	Label []Analysis_DimensionField `json:"Label,omitempty"`

	// Size AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells-size
	Size []Analysis_MeasureField `json:"Size,omitempty"`

	// XAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells-xaxis
	XAxis []Analysis_MeasureField `json:"XAxis,omitempty"`

	// YAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-analysis-scatterplotcategoricallyaggregatedfieldwells-yaxis
	YAxis []Analysis_MeasureField `json:"YAxis,omitempty"`

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
func (r *Analysis_ScatterPlotCategoricallyAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ScatterPlotCategoricallyAggregatedFieldWells"
}