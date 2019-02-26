package cloudformation

// AWSKinesisAnalyticsV2Application_InputParallelism AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.InputParallelism)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputparallelism.html
type AWSKinesisAnalyticsV2Application_InputParallelism struct {

	// Count AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputparallelism.html#cfn-kinesisanalyticsv2-application-inputparallelism-count
	Count int `json:"Count,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_InputParallelism) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.InputParallelism"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_InputParallelism) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
