// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_PivotTableCellConditionalFormatting AWS CloudFormation Resource (AWS::QuickSight::Template.PivotTableCellConditionalFormatting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablecellconditionalformatting.html
type Template_PivotTableCellConditionalFormatting struct {

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablecellconditionalformatting.html#cfn-quicksight-template-pivottablecellconditionalformatting-fieldid
	FieldId string `json:"FieldId"`

	// Scope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablecellconditionalformatting.html#cfn-quicksight-template-pivottablecellconditionalformatting-scope
	Scope *Template_PivotTableConditionalFormattingScope `json:"Scope,omitempty"`

	// Scopes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablecellconditionalformatting.html#cfn-quicksight-template-pivottablecellconditionalformatting-scopes
	Scopes []Template_PivotTableConditionalFormattingScope `json:"Scopes,omitempty"`

	// TextFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablecellconditionalformatting.html#cfn-quicksight-template-pivottablecellconditionalformatting-textformat
	TextFormat *Template_TextConditionalFormat `json:"TextFormat,omitempty"`

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
func (r *Template_PivotTableCellConditionalFormatting) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PivotTableCellConditionalFormatting"
}
