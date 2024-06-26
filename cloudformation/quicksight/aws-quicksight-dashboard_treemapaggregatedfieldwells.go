// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_TreeMapAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TreeMapAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-treemapaggregatedfieldwells.html
type Dashboard_TreeMapAggregatedFieldWells struct {

	// Colors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-treemapaggregatedfieldwells.html#cfn-quicksight-dashboard-treemapaggregatedfieldwells-colors
	Colors []Dashboard_MeasureField `json:"Colors,omitempty"`

	// Groups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-treemapaggregatedfieldwells.html#cfn-quicksight-dashboard-treemapaggregatedfieldwells-groups
	Groups []Dashboard_DimensionField `json:"Groups,omitempty"`

	// Sizes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-treemapaggregatedfieldwells.html#cfn-quicksight-dashboard-treemapaggregatedfieldwells-sizes
	Sizes []Dashboard_MeasureField `json:"Sizes,omitempty"`

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
func (r *Dashboard_TreeMapAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TreeMapAggregatedFieldWells"
}
