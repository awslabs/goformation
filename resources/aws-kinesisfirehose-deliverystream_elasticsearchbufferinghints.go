package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html
type AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints struct {

	// IntervalInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints-intervalinseconds
	IntervalInSeconds int64 `json:"IntervalInSeconds"`

	// SizeInMBs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-bufferinghints-sizeinmbs
	SizeInMBs int64 `json:"SizeInMBs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHintsResources retrieves all AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHintsWithName retrieves all AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints, error) {

	result := &AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints{}, errors.New("resource not found")

}
