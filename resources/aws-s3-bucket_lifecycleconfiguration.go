package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.LifecycleConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
type AWSS3Bucket_LifecycleConfiguration struct {

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html#cfn-s3-bucket-lifecycleconfig-rules
	Rules []AWSS3Bucket_Rule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_LifecycleConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.LifecycleConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_LifecycleConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_LifecycleConfigurationResources retrieves all AWSS3Bucket_LifecycleConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_LifecycleConfiguration(template *Template) map[string]*AWSS3Bucket_LifecycleConfiguration {

	results := map[string]*AWSS3Bucket_LifecycleConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_LifecycleConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_LifecycleConfigurationWithName retrieves all AWSS3Bucket_LifecycleConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_LifecycleConfiguration(name string, template *Template) (*AWSS3Bucket_LifecycleConfiguration, error) {

	result := &AWSS3Bucket_LifecycleConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_LifecycleConfiguration{}, errors.New("resource not found")

}
