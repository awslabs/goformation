// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_SankeyDiagramFieldWells AWS CloudFormation Resource (AWS::QuickSight::Template.SankeyDiagramFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramfieldwells.html
type Template_SankeyDiagramFieldWells struct {

	// SankeyDiagramAggregatedFieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramfieldwells.html#cfn-quicksight-template-sankeydiagramfieldwells-sankeydiagramaggregatedfieldwells
	SankeyDiagramAggregatedFieldWells *Template_SankeyDiagramAggregatedFieldWells `json:"SankeyDiagramAggregatedFieldWells,omitempty"`

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
func (r *Template_SankeyDiagramFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.SankeyDiagramFieldWells"
}
