package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.CorsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
type AWSS3Bucket_CorsConfiguration struct {

	// CorsRules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html#cfn-s3-bucket-cors-corsrule
	CorsRules []AWSS3Bucket_CorsRule `json:"CorsRules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_CorsConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.CorsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_CorsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_CorsConfigurationResources retrieves all AWSS3Bucket_CorsConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_CorsConfiguration(template *Template) map[string]*AWSS3Bucket_CorsConfiguration {

	results := map[string]*AWSS3Bucket_CorsConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_CorsConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_CorsConfigurationWithName retrieves all AWSS3Bucket_CorsConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_CorsConfiguration(name string, template *Template) (*AWSS3Bucket_CorsConfiguration, error) {

	result := &AWSS3Bucket_CorsConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_CorsConfiguration{}, errors.New("resource not found")

}
