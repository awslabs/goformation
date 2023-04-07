// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_CategoricalDimensionField AWS CloudFormation Resource (AWS::QuickSight::Template.CategoricalDimensionField)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoricaldimensionfield.html
type Template_CategoricalDimensionField struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoricaldimensionfield.html#cfn-quicksight-template-categoricaldimensionfield-column
	Column *Template_ColumnIdentifier `json:"Column"`

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoricaldimensionfield.html#cfn-quicksight-template-categoricaldimensionfield-fieldid
	FieldId string `json:"FieldId"`

	// FormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoricaldimensionfield.html#cfn-quicksight-template-categoricaldimensionfield-formatconfiguration
	FormatConfiguration *Template_StringFormatConfiguration `json:"FormatConfiguration,omitempty"`

	// HierarchyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoricaldimensionfield.html#cfn-quicksight-template-categoricaldimensionfield-hierarchyid
	HierarchyId *string `json:"HierarchyId,omitempty"`

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
func (r *Template_CategoricalDimensionField) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.CategoricalDimensionField"
}