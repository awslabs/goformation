// Code generated by "go generate". Please don't change this file directly.

package ssmcontacts

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Plan_ContactTargetInfo AWS CloudFormation Resource (AWS::SSMContacts::Plan.ContactTargetInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-plan-contacttargetinfo.html
type Plan_ContactTargetInfo struct {

	// ContactId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-plan-contacttargetinfo.html#cfn-ssmcontacts-plan-contacttargetinfo-contactid
	ContactId string `json:"ContactId"`

	// IsEssential AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-plan-contacttargetinfo.html#cfn-ssmcontacts-plan-contacttargetinfo-isessential
	IsEssential bool `json:"IsEssential"`

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
func (r *Plan_ContactTargetInfo) AWSCloudFormationType() string {
	return "AWS::SSMContacts::Plan.ContactTargetInfo"
}