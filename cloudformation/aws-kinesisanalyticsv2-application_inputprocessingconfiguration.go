package cloudformation

// AWSKinesisAnalyticsV2Application_InputProcessingConfiguration AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.InputProcessingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputprocessingconfiguration.html
type AWSKinesisAnalyticsV2Application_InputProcessingConfiguration struct {

	// InputLambdaProcessor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputprocessingconfiguration.html#cfn-kinesisanalyticsv2-application-inputprocessingconfiguration-inputlambdaprocessor
	InputLambdaProcessor *AWSKinesisAnalyticsV2Application_InputLambdaProcessor `json:"InputLambdaProcessor,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_InputProcessingConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.InputProcessingConfiguration"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_InputProcessingConfiguration) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
