// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_TooltipOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.TooltipOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html
type Analysis_TooltipOptions struct {

	// FieldBasedTooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html#cfn-quicksight-analysis-tooltipoptions-fieldbasedtooltip
	FieldBasedTooltip *Analysis_FieldBasedTooltip `json:"FieldBasedTooltip,omitempty"`

	// SelectedTooltipType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html#cfn-quicksight-analysis-tooltipoptions-selectedtooltiptype
	SelectedTooltipType *string `json:"SelectedTooltipType,omitempty"`

	// TooltipVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tooltipoptions.html#cfn-quicksight-analysis-tooltipoptions-tooltipvisibility
	TooltipVisibility *string `json:"TooltipVisibility,omitempty"`

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
func (r *Analysis_TooltipOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TooltipOptions"
}