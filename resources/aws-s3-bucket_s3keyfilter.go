package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.S3KeyFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
type AWSS3Bucket_S3KeyFilter struct {

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules

	Rules []AWSS3Bucket_FilterRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_S3KeyFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.S3KeyFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_S3KeyFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_S3KeyFilterResources retrieves all AWSS3Bucket_S3KeyFilter items from a CloudFormation template
func GetAllAWSS3Bucket_S3KeyFilter(template *Template) map[string]*AWSS3Bucket_S3KeyFilter {

	results := map[string]*AWSS3Bucket_S3KeyFilter{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_S3KeyFilter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_S3KeyFilterWithName retrieves all AWSS3Bucket_S3KeyFilter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_S3KeyFilter(name string, template *Template) (*AWSS3Bucket_S3KeyFilter, error) {

	result := &AWSS3Bucket_S3KeyFilter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_S3KeyFilter{}, errors.New("resource not found")

}
