package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html
type AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions-enabled
	Enabled bool `json:"Enabled"`

	// LogGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions-loggroupname
	LogGroupName string `json:"LogGroupName"`

	// LogStreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-destination-cloudwatchloggingoptions-logstreamname
	LogStreamName string `json:"LogStreamName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.CloudWatchLoggingOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
