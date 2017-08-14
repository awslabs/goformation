package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.NotificationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html
type AWSS3Bucket_NotificationConfiguration struct {

	// LambdaConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig

	LambdaConfigurations []AWSS3Bucket_LambdaConfiguration `json:"LambdaConfigurations"`

	// QueueConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-queueconfig

	QueueConfigurations []AWSS3Bucket_QueueConfiguration `json:"QueueConfigurations"`

	// TopicConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-topicconfig

	TopicConfigurations []AWSS3Bucket_TopicConfiguration `json:"TopicConfigurations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_NotificationConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.NotificationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_NotificationConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_NotificationConfigurationResources retrieves all AWSS3Bucket_NotificationConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_NotificationConfiguration(template *Template) map[string]*AWSS3Bucket_NotificationConfiguration {

	results := map[string]*AWSS3Bucket_NotificationConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_NotificationConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_NotificationConfigurationWithName retrieves all AWSS3Bucket_NotificationConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_NotificationConfiguration(name string, template *Template) (*AWSS3Bucket_NotificationConfiguration, error) {

	result := &AWSS3Bucket_NotificationConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_NotificationConfiguration{}, errors.New("resource not found")

}
