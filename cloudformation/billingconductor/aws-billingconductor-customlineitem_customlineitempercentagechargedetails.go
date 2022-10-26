// Code generated by "go generate". Please don't change this file directly.

package billingconductor

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// CustomLineItem_CustomLineItemPercentageChargeDetails AWS CloudFormation Resource (AWS::BillingConductor::CustomLineItem.CustomLineItemPercentageChargeDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html
type CustomLineItem_CustomLineItemPercentageChargeDetails struct {

	// ChildAssociatedResources AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html#cfn-billingconductor-customlineitem-customlineitempercentagechargedetails-childassociatedresources
	ChildAssociatedResources []string `json:"ChildAssociatedResources,omitempty"`

	// PercentageValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html#cfn-billingconductor-customlineitem-customlineitempercentagechargedetails-percentagevalue
	PercentageValue float64 `json:"PercentageValue"`

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
func (r *CustomLineItem_CustomLineItemPercentageChargeDetails) AWSCloudFormationType() string {
	return "AWS::BillingConductor::CustomLineItem.CustomLineItemPercentageChargeDetails"
}