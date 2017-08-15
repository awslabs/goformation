package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.NoncurrentVersionTransition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html
type AWSS3Bucket_NoncurrentVersionTransition struct {

	// StorageClass AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-storageclass

	StorageClass string `json:"StorageClass"`

	// TransitionInDays AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition-transitionindays

	TransitionInDays int64 `json:"TransitionInDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_NoncurrentVersionTransition) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.NoncurrentVersionTransition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_NoncurrentVersionTransition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_NoncurrentVersionTransitionResources retrieves all AWSS3Bucket_NoncurrentVersionTransition items from a CloudFormation template
func GetAllAWSS3Bucket_NoncurrentVersionTransition(template *Template) map[string]*AWSS3Bucket_NoncurrentVersionTransition {

	results := map[string]*AWSS3Bucket_NoncurrentVersionTransition{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_NoncurrentVersionTransition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_NoncurrentVersionTransitionWithName retrieves all AWSS3Bucket_NoncurrentVersionTransition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_NoncurrentVersionTransition(name string, template *Template) (*AWSS3Bucket_NoncurrentVersionTransition, error) {

	result := &AWSS3Bucket_NoncurrentVersionTransition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_NoncurrentVersionTransition{}, errors.New("resource not found")

}
