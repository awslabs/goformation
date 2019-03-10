package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSKinesisAnalyticsV2Application_EnvironmentProperties AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.EnvironmentProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html
type AWSKinesisAnalyticsV2Application_EnvironmentProperties struct {

	// PropertyGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html#cfn-kinesisanalyticsv2-application-environmentproperties-propertygroups
	PropertyGroups []AWSKinesisAnalyticsV2Application_PropertyGroup `json:"PropertyGroups,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.EnvironmentProperties"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
