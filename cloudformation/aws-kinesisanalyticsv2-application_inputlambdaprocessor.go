package cloudformation

// AWSKinesisAnalyticsV2Application_InputLambdaProcessor AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.InputLambdaProcessor)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputlambdaprocessor.html
type AWSKinesisAnalyticsV2Application_InputLambdaProcessor struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputlambdaprocessor.html#cfn-kinesisanalyticsv2-application-inputlambdaprocessor-resourcearn
	ResourceARN string `json:"ResourceARN,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_InputLambdaProcessor) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.InputLambdaProcessor"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_InputLambdaProcessor) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
