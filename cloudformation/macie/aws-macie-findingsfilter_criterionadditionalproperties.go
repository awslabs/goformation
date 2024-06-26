// Code generated by "go generate". Please don't change this file directly.

package macie

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// FindingsFilter_CriterionAdditionalProperties AWS CloudFormation Resource (AWS::Macie::FindingsFilter.CriterionAdditionalProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html
type FindingsFilter_CriterionAdditionalProperties struct {

	// eq AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-eq
	eq []string `json:"eq,omitempty"`

	// gt AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-gt
	gt *int `json:"gt,omitempty"`

	// gte AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-gte
	gte *int `json:"gte,omitempty"`

	// lt AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-lt
	lt *int `json:"lt,omitempty"`

	// lte AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-lte
	lte *int `json:"lte,omitempty"`

	// neq AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-neq
	neq []string `json:"neq,omitempty"`

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
func (r *FindingsFilter_CriterionAdditionalProperties) AWSCloudFormationType() string {
	return "AWS::Macie::FindingsFilter.CriterionAdditionalProperties"
}
