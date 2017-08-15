package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions.html
type AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions struct {

	// DurationInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions.html#cfn-kinesisfirehose-kinesisdeliverystream-elasticsearchdestinationconfiguration-retryoptions-durationinseconds
	DurationInSeconds int64 `json:"DurationInSeconds"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.ElasticsearchRetryOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptionsResources retrieves all AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptionsWithName retrieves all AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions, error) {

	result := &AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_ElasticsearchRetryOptions{}, errors.New("resource not found")

}
