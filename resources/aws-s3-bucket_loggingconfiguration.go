package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.LoggingConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html
type AWSS3Bucket_LoggingConfiguration struct {

	// DestinationBucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-destinationbucketname
	DestinationBucketName string `json:"DestinationBucketName"`

	// LogFilePrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-loggingconfig.html#cfn-s3-bucket-loggingconfig-logfileprefix
	LogFilePrefix string `json:"LogFilePrefix"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_LoggingConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.LoggingConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_LoggingConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_LoggingConfigurationResources retrieves all AWSS3Bucket_LoggingConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_LoggingConfiguration(template *Template) map[string]*AWSS3Bucket_LoggingConfiguration {

	results := map[string]*AWSS3Bucket_LoggingConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_LoggingConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_LoggingConfigurationWithName retrieves all AWSS3Bucket_LoggingConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_LoggingConfiguration(name string, template *Template) (*AWSS3Bucket_LoggingConfiguration, error) {

	result := &AWSS3Bucket_LoggingConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_LoggingConfiguration{}, errors.New("resource not found")

}
