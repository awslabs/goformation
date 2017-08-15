package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html
type AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration struct {

	// KMSEncryptionConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig

	KMSEncryptionConfig AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig `json:"KMSEncryptionConfig"`

	// NoEncryptionConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-noencryptionconfig

	NoEncryptionConfig string `json:"NoEncryptionConfig"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_EncryptionConfigurationResources retrieves all AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_EncryptionConfiguration(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_EncryptionConfigurationWithName retrieves all AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_EncryptionConfiguration(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration, error) {

	result := &AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration{}, errors.New("resource not found")

}
