package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.QueueConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html
type AWSS3Bucket_QueueConfiguration struct {

	// Event AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-event

	Event string `json:"Event"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-filter

	Filter AWSS3Bucket_NotificationFilter `json:"Filter"`

	// Queue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-queueconfig.html#cfn-s3-bucket-notificationconfig-queueconfig-queue

	Queue string `json:"Queue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_QueueConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.QueueConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_QueueConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_QueueConfigurationResources retrieves all AWSS3Bucket_QueueConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_QueueConfiguration(template *Template) map[string]*AWSS3Bucket_QueueConfiguration {

	results := map[string]*AWSS3Bucket_QueueConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_QueueConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_QueueConfigurationWithName retrieves all AWSS3Bucket_QueueConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_QueueConfiguration(name string, template *Template) (*AWSS3Bucket_QueueConfiguration, error) {

	result := &AWSS3Bucket_QueueConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_QueueConfiguration{}, errors.New("resource not found")

}
