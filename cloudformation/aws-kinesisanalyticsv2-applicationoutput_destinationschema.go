package cloudformation

// AWSKinesisAnalyticsV2ApplicationOutput_DestinationSchema AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::ApplicationOutput.DestinationSchema)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-destinationschema.html
type AWSKinesisAnalyticsV2ApplicationOutput_DestinationSchema struct {

	// RecordFormatType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-destinationschema.html#cfn-kinesisanalyticsv2-applicationoutput-destinationschema-recordformattype
	RecordFormatType string `json:"RecordFormatType,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2ApplicationOutput_DestinationSchema) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::ApplicationOutput.DestinationSchema"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2ApplicationOutput_DestinationSchema) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
