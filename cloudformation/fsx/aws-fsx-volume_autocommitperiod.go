// Code generated by "go generate". Please don't change this file directly.

package fsx

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Volume_AutocommitPeriod AWS CloudFormation Resource (AWS::FSx::Volume.AutocommitPeriod)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration-autocommitperiod.html
type Volume_AutocommitPeriod struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration-autocommitperiod.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-autocommitperiod-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration-autocommitperiod.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-autocommitperiod-value
	Value *int `json:"Value,omitempty"`

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
func (r *Volume_AutocommitPeriod) AWSCloudFormationType() string {
	return "AWS::FSx::Volume.AutocommitPeriod"
}
