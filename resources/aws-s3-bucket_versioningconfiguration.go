package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.VersioningConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
type AWSS3Bucket_VersioningConfiguration struct {

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html#cfn-s3-bucket-versioningconfig-status

	Status string `json:"Status"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_VersioningConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.VersioningConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_VersioningConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_VersioningConfigurationResources retrieves all AWSS3Bucket_VersioningConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_VersioningConfiguration(template *Template) map[string]*AWSS3Bucket_VersioningConfiguration {

	results := map[string]*AWSS3Bucket_VersioningConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_VersioningConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_VersioningConfigurationWithName retrieves all AWSS3Bucket_VersioningConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_VersioningConfiguration(name string, template *Template) (*AWSS3Bucket_VersioningConfiguration, error) {

	result := &AWSS3Bucket_VersioningConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_VersioningConfiguration{}, errors.New("resource not found")

}
