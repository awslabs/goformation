package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.NotificationFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
type AWSS3Bucket_NotificationFilter struct {

	// S3Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key

	S3Key AWSS3Bucket_S3KeyFilter `json:"S3Key"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_NotificationFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.NotificationFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_NotificationFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_NotificationFilterResources retrieves all AWSS3Bucket_NotificationFilter items from a CloudFormation template
func GetAllAWSS3Bucket_NotificationFilter(template *Template) map[string]*AWSS3Bucket_NotificationFilter {

	results := map[string]*AWSS3Bucket_NotificationFilter{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_NotificationFilter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_NotificationFilterWithName retrieves all AWSS3Bucket_NotificationFilter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_NotificationFilter(name string, template *Template) (*AWSS3Bucket_NotificationFilter, error) {

	result := &AWSS3Bucket_NotificationFilter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_NotificationFilter{}, errors.New("resource not found")

}
