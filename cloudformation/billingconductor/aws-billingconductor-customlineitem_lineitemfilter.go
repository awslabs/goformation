// Code generated by "go generate". Please don't change this file directly.

package billingconductor

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// CustomLineItem_LineItemFilter AWS CloudFormation Resource (AWS::BillingConductor::CustomLineItem.LineItemFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html
type CustomLineItem_LineItemFilter struct {

	// Attribute AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html#cfn-billingconductor-customlineitem-lineitemfilter-attribute
	Attribute string `json:"Attribute"`

	// MatchOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html#cfn-billingconductor-customlineitem-lineitemfilter-matchoption
	MatchOption string `json:"MatchOption"`

	// Values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-lineitemfilter.html#cfn-billingconductor-customlineitem-lineitemfilter-values
	Values []string `json:"Values"`

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
func (r *CustomLineItem_LineItemFilter) AWSCloudFormationType() string {
	return "AWS::BillingConductor::CustomLineItem.LineItemFilter"
}