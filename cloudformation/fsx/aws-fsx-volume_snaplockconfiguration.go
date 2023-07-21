// Code generated by "go generate". Please don't change this file directly.

package fsx

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Volume_SnaplockConfiguration AWS CloudFormation Resource (AWS::FSx::Volume.SnaplockConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html
type Volume_SnaplockConfiguration struct {

	// AuditLogVolume AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-auditlogvolume
	AuditLogVolume *string `json:"AuditLogVolume,omitempty"`

	// AutocommitPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-autocommitperiod
	AutocommitPeriod *Volume_AutocommitPeriod `json:"AutocommitPeriod,omitempty"`

	// PrivilegedDelete AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-privilegeddelete
	PrivilegedDelete *string `json:"PrivilegedDelete,omitempty"`

	// RetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-retentionperiod
	RetentionPeriod *Volume_SnaplockRetentionPeriod `json:"RetentionPeriod,omitempty"`

	// SnaplockType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-snaplocktype
	SnaplockType string `json:"SnaplockType"`

	// VolumeAppendModeEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration-snaplockconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration-volumeappendmodeenabled
	VolumeAppendModeEnabled *string `json:"VolumeAppendModeEnabled,omitempty"`

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
func (r *Volume_SnaplockConfiguration) AWSCloudFormationType() string {
	return "AWS::FSx::Volume.SnaplockConfiguration"
}
