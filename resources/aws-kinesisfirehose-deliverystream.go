package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
type AWSKinesisFirehoseDeliveryStream struct {

	// DeliveryStreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverstream-deliverystreamname

	DeliveryStreamName string `json:"DeliveryStreamName"`

	// ElasticsearchDestinationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverstream-elasticsearchdestinationconfiguration

	ElasticsearchDestinationConfiguration AWSKinesisFirehoseDeliveryStream_ElasticsearchDestinationConfiguration `json:"ElasticsearchDestinationConfiguration"`

	// RedshiftDestinationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration

	RedshiftDestinationConfiguration AWSKinesisFirehoseDeliveryStream_RedshiftDestinationConfiguration `json:"RedshiftDestinationConfiguration"`

	// S3DestinationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration

	S3DestinationConfiguration AWSKinesisFirehoseDeliveryStream_S3DestinationConfiguration `json:"S3DestinationConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStreamResources retrieves all AWSKinesisFirehoseDeliveryStream items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream {

	results := map[string]*AWSKinesisFirehoseDeliveryStream{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStreamWithName retrieves all AWSKinesisFirehoseDeliveryStream items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream, error) {

	result := &AWSKinesisFirehoseDeliveryStream{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream{}, errors.New("resource not found")

}
