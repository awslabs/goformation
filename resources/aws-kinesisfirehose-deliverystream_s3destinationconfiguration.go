package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html
type AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration struct {

	// BucketARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bucketarn
	BucketARN string `json:"BucketARN"`

	// BufferingHints AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints
	BufferingHints AWSKinesisFirehoseDeliveryStream_BufferingHints `json:"BufferingHints"`

	// CloudWatchLoggingOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-cloudwatchloggingoptions
	CloudWatchLoggingOptions AWSKinesisFirehoseDeliveryStream_CloudWatchLoggingOptions `json:"CloudWatchLoggingOptions"`

	// CompressionFormat AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-compressionformat
	CompressionFormat string `json:"CompressionFormat"`

	// EncryptionConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration
	EncryptionConfiguration AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration `json:"EncryptionConfiguration"`

	// Prefix AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-prefix
	Prefix string `json:"Prefix"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-rolearn
	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_S3DestinationConfigurationResources retrieves all AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_S3DestinationConfigurationWithName retrieves all AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration, error) {

	result := &AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration{}, errors.New("resource not found")

}
