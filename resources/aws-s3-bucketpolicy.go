package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::BucketPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
type AWSS3BucketPolicy struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html#cfn-s3-bucketpolicy-bucket

	Bucket string `json:"Bucket"`

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html#cfn-s3-bucketpolicy-policydocument

	PolicyDocument interface{} `json:"PolicyDocument"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketPolicy) AWSCloudFormationType() string {
	return "AWS::S3::BucketPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3BucketPolicyResources retrieves all AWSS3BucketPolicy items from a CloudFormation template
func GetAllAWSS3BucketPolicy(template *Template) map[string]*AWSS3BucketPolicy {

	results := map[string]*AWSS3BucketPolicy{}
	for name, resource := range template.Resources {
		result := &AWSS3BucketPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3BucketPolicyWithName retrieves all AWSS3BucketPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3BucketPolicy(name string, template *Template) (*AWSS3BucketPolicy, error) {

	result := &AWSS3BucketPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3BucketPolicy{}, errors.New("resource not found")

}
