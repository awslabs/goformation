// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_ScatterPlotUnaggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ScatterPlotUnaggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html
type Dashboard_ScatterPlotUnaggregatedFieldWells struct {

	// Category AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-category
	Category []Dashboard_DimensionField `json:"Category,omitempty"`

	// Label AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-label
	Label []Dashboard_DimensionField `json:"Label,omitempty"`

	// Size AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-size
	Size []Dashboard_MeasureField `json:"Size,omitempty"`

	// XAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-xaxis
	XAxis []Dashboard_DimensionField `json:"XAxis,omitempty"`

	// YAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-yaxis
	YAxis []Dashboard_DimensionField `json:"YAxis,omitempty"`

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
func (r *Dashboard_ScatterPlotUnaggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ScatterPlotUnaggregatedFieldWells"
}
