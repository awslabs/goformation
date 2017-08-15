package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.ReplicationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html
type AWSS3Bucket_ReplicationConfiguration struct {

	// Role AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-role
	Role string `json:"Role"`

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-rules
	Rules []AWSS3Bucket_ReplicationRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_ReplicationConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.ReplicationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_ReplicationConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_ReplicationConfigurationResources retrieves all AWSS3Bucket_ReplicationConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_ReplicationConfiguration(template *Template) map[string]*AWSS3Bucket_ReplicationConfiguration {

	results := map[string]*AWSS3Bucket_ReplicationConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_ReplicationConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_ReplicationConfigurationWithName retrieves all AWSS3Bucket_ReplicationConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_ReplicationConfiguration(name string, template *Template) (*AWSS3Bucket_ReplicationConfiguration, error) {

	result := &AWSS3Bucket_ReplicationConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_ReplicationConfiguration{}, errors.New("resource not found")

}
