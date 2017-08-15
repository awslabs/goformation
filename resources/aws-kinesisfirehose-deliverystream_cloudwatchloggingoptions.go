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

// GetAllAWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptionsResources retrieves all AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptionsWithName retrieves all AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions, error) {

	result := &AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions{}, errors.New("resource not found")

}
