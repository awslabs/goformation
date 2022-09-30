// Code generated by "go generate". Please don't change this file directly.

package imagebuilder

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DistributionConfiguration_AmiDistributionConfiguration AWS CloudFormation Resource (AWS::ImageBuilder::DistributionConfiguration.AmiDistributionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html
type DistributionConfiguration_AmiDistributionConfiguration struct {

	// AmiTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-amitags
	AmiTags map[string]string `json:"AmiTags,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-description
	Description *string `json:"Description,omitempty"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-kmskeyid
	KmsKeyId *string `json:"KmsKeyId,omitempty"`

	// LaunchPermissionConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-launchpermissionconfiguration
	LaunchPermissionConfiguration *DistributionConfiguration_LaunchPermissionConfiguration `json:"LaunchPermissionConfiguration,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-name
	Name *string `json:"Name,omitempty"`

	// TargetAccountIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-amidistributionconfiguration-targetaccountids
	TargetAccountIds []string `json:"TargetAccountIds,omitempty"`

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
func (r *DistributionConfiguration_AmiDistributionConfiguration) AWSCloudFormationType() string {
	return "AWS::ImageBuilder::DistributionConfiguration.AmiDistributionConfiguration"
}