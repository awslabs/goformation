package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.LambdaConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html
type AWSS3Bucket_LambdaConfiguration struct {

	// Event AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-event
	Event string `json:"Event"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-filter
	Filter AWSS3Bucket_NotificationFilter `json:"Filter"`

	// Function AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig-lambdaconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig-function
	Function string `json:"Function"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_LambdaConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.LambdaConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_LambdaConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_LambdaConfigurationResources retrieves all AWSS3Bucket_LambdaConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_LambdaConfiguration(template *Template) map[string]*AWSS3Bucket_LambdaConfiguration {

	results := map[string]*AWSS3Bucket_LambdaConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_LambdaConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_LambdaConfigurationWithName retrieves all AWSS3Bucket_LambdaConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_LambdaConfiguration(name string, template *Template) (*AWSS3Bucket_LambdaConfiguration, error) {

	result := &AWSS3Bucket_LambdaConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_LambdaConfiguration{}, errors.New("resource not found")

}
