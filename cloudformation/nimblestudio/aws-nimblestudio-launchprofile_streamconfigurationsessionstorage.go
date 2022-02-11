package nimblestudio

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// LaunchProfile_StreamConfigurationSessionStorage AWS CloudFormation Resource (AWS::NimbleStudio::LaunchProfile.StreamConfigurationSessionStorage)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionstorage.html
type LaunchProfile_StreamConfigurationSessionStorage struct {

	// Mode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionstorage.html#cfn-nimblestudio-launchprofile-streamconfigurationsessionstorage-mode
	Mode *[]string `json:"Mode,omitempty"`

	// Root AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfigurationsessionstorage.html#cfn-nimblestudio-launchprofile-streamconfigurationsessionstorage-root
	Root *LaunchProfile_StreamingSessionStorageRoot `json:"Root,omitempty"`

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
func (r *LaunchProfile_StreamConfigurationSessionStorage) AWSCloudFormationType() string {
	return "AWS::NimbleStudio::LaunchProfile.StreamConfigurationSessionStorage"
}
