// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_DefaultFilterDropDownControlOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.DefaultFilterDropDownControlOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfilterdropdowncontroloptions.html
type Analysis_DefaultFilterDropDownControlOptions struct {

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfilterdropdowncontroloptions.html#cfn-quicksight-analysis-defaultfilterdropdowncontroloptions-displayoptions
	DisplayOptions *Analysis_DropDownControlDisplayOptions `json:"DisplayOptions,omitempty"`

	// SelectableValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfilterdropdowncontroloptions.html#cfn-quicksight-analysis-defaultfilterdropdowncontroloptions-selectablevalues
	SelectableValues *Analysis_FilterSelectableValues `json:"SelectableValues,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfilterdropdowncontroloptions.html#cfn-quicksight-analysis-defaultfilterdropdowncontroloptions-type
	Type *string `json:"Type,omitempty"`

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
func (r *Analysis_DefaultFilterDropDownControlOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DefaultFilterDropDownControlOptions"
}
