package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSMediaLiveChannel_VideoSelector AWS CloudFormation Resource (AWS::MediaLive::Channel.VideoSelector)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html
type AWSMediaLiveChannel_VideoSelector struct {

	// ColorSpace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-colorspace
	ColorSpace string `json:"ColorSpace,omitempty"`

	// ColorSpaceUsage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-colorspaceusage
	ColorSpaceUsage string `json:"ColorSpaceUsage,omitempty"`

	// SelectorSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselector.html#cfn-medialive-channel-videoselector-selectorsettings
	SelectorSettings *AWSMediaLiveChannel_VideoSelectorSettings `json:"SelectorSettings,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSMediaLiveChannel_VideoSelector) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.VideoSelector"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSMediaLiveChannel_VideoSelector) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSMediaLiveChannel_VideoSelector) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSMediaLiveChannel_VideoSelector) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSMediaLiveChannel_VideoSelector) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSMediaLiveChannel_VideoSelector) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
