package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig.html
type AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig struct {

	// AWSKMSKeyARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-encryptionconfiguration-kmsencryptionconfig-awskmskeyarn
	AWSKMSKeyARN string `json:"AWSKMSKeyARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisFirehoseDeliveryStream_KMSEncryptionConfigResources retrieves all AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig items from a CloudFormation template
func GetAllAWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig(template *Template) map[string]*AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig {

	results := map[string]*AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig{}
	for name, resource := range template.Resources {
		result := &AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisFirehoseDeliveryStream_KMSEncryptionConfigWithName retrieves all AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig(name string, template *Template) (*AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig, error) {

	result := &AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig{}, errors.New("resource not found")

}
