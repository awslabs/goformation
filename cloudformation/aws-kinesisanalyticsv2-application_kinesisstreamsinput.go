package cloudformation

// AWSKinesisAnalyticsV2Application_KinesisStreamsInput AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.KinesisStreamsInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-kinesisstreamsinput.html
type AWSKinesisAnalyticsV2Application_KinesisStreamsInput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-kinesisstreamsinput.html#cfn-kinesisanalyticsv2-application-kinesisstreamsinput-resourcearn
	ResourceARN string `json:"ResourceARN,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_KinesisStreamsInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.KinesisStreamsInput"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_KinesisStreamsInput) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
