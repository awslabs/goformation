package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSKinesisAnalyticsV2Application_KinesisFirehoseInput AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.KinesisFirehoseInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-kinesisfirehoseinput.html
type AWSKinesisAnalyticsV2Application_KinesisFirehoseInput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-kinesisfirehoseinput.html#cfn-kinesisanalyticsv2-application-kinesisfirehoseinput-resourcearn
	ResourceARN string `json:"ResourceARN,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_KinesisFirehoseInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.KinesisFirehoseInput"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSKinesisAnalyticsV2Application_KinesisFirehoseInput) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSKinesisAnalyticsV2Application_KinesisFirehoseInput) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSKinesisAnalyticsV2Application_KinesisFirehoseInput) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSKinesisAnalyticsV2Application_KinesisFirehoseInput) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_KinesisFirehoseInput) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
