// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_KPIActualValueConditionalFormatting AWS CloudFormation Resource (AWS::QuickSight::Template.KPIActualValueConditionalFormatting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiactualvalueconditionalformatting.html
type Template_KPIActualValueConditionalFormatting struct {

	// Icon AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiactualvalueconditionalformatting.html#cfn-quicksight-template-kpiactualvalueconditionalformatting-icon
	Icon *Template_ConditionalFormattingIcon `json:"Icon,omitempty"`

	// TextColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiactualvalueconditionalformatting.html#cfn-quicksight-template-kpiactualvalueconditionalformatting-textcolor
	TextColor *Template_ConditionalFormattingColor `json:"TextColor,omitempty"`

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
func (r *Template_KPIActualValueConditionalFormatting) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.KPIActualValueConditionalFormatting"
}
