package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSPinpointApplicationSettings_Limits AWS CloudFormation Resource (AWS::Pinpoint::ApplicationSettings.Limits)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html
type AWSPinpointApplicationSettings_Limits struct {

	// Daily AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-daily
	Daily int `json:"Daily,omitempty"`

	// MaximumDuration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-maximumduration
	MaximumDuration int `json:"MaximumDuration,omitempty"`

	// MessagesPerSecond AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-messagespersecond
	MessagesPerSecond int `json:"MessagesPerSecond,omitempty"`

	// Total AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-limits.html#cfn-pinpoint-applicationsettings-limits-total
	Total int `json:"Total,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSPinpointApplicationSettings_Limits) AWSCloudFormationType() string {
	return "AWS::Pinpoint::ApplicationSettings.Limits"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSPinpointApplicationSettings_Limits) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSPinpointApplicationSettings_Limits) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSPinpointApplicationSettings_Limits) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSPinpointApplicationSettings_Limits) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSPinpointApplicationSettings_Limits) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
