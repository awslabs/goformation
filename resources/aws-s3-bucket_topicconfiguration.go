package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.TopicConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html
type AWSS3Bucket_TopicConfiguration struct {

	// Event AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-event
	Event string `json:"Event"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-filter
	Filter AWSS3Bucket_NotificationFilter `json:"Filter"`

	// Topic AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-topicconfig.html#cfn-s3-bucket-notificationconfig-topicconfig-topic
	Topic string `json:"Topic"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_TopicConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.TopicConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_TopicConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_TopicConfigurationResources retrieves all AWSS3Bucket_TopicConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_TopicConfiguration(template *Template) map[string]*AWSS3Bucket_TopicConfiguration {

	results := map[string]*AWSS3Bucket_TopicConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_TopicConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_TopicConfigurationWithName retrieves all AWSS3Bucket_TopicConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_TopicConfiguration(name string, template *Template) (*AWSS3Bucket_TopicConfiguration, error) {

	result := &AWSS3Bucket_TopicConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_TopicConfiguration{}, errors.New("resource not found")

}
