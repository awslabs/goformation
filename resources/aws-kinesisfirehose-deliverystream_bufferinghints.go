package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.BufferingHints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html
type AWSKinesisFirehoseDeliveryStream_BufferingHints struct {

	// IntervalInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints-intervalinseconds
	IntervalInSeconds int64 `json:"IntervalInSeconds"`

	// SizeInMBs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints-sizeinmbs
	SizeInMBs int64 `json:"SizeInMBs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_BufferingHints) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.BufferingHints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_BufferingHints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_BufferingHintsResources retrieves all AWSKinesisFirehoseDeliveryStream_BufferingHints items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_BufferingHints(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_BufferingHints {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_BufferingHints{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_BufferingHints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_BufferingHintsWithName retrieves all AWSKinesisFirehoseDeliveryStream_BufferingHints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_BufferingHints(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_BufferingHints, error) {

	result := &AWSKinesisFirehoseDeliveryStream_BufferingHints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_BufferingHints{}, errors.New("resource not found")

}
