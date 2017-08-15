package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSS3Bucket AWS CloudFormation Resource (AWS::S3::Bucket)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
type AWSS3Bucket struct {

	// AccessControl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
	AccessControl string `json:"AccessControl"`

	// BucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
	BucketName string `json:"BucketName"`

	// CorsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
	CorsConfiguration AWSS3Bucket_CorsConfiguration `json:"CorsConfiguration"`

	// LifecycleConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
	LifecycleConfiguration AWSS3Bucket_LifecycleConfiguration `json:"LifecycleConfiguration"`

	// LoggingConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
	LoggingConfiguration AWSS3Bucket_LoggingConfiguration `json:"LoggingConfiguration"`

	// NotificationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
	NotificationConfiguration AWSS3Bucket_NotificationConfiguration `json:"NotificationConfiguration"`

	// ReplicationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
	ReplicationConfiguration AWSS3Bucket_ReplicationConfiguration `json:"ReplicationConfiguration"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags
	Tags []Tag `json:"Tags"`

	// VersioningConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
	VersioningConfiguration AWSS3Bucket_VersioningConfiguration `json:"VersioningConfiguration"`

	// WebsiteConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
	WebsiteConfiguration AWSS3Bucket_WebsiteConfiguration `json:"WebsiteConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket) AWSCloudFormationType() string {
	return "AWS::S3::Bucket"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3BucketResources retrieves all AWSS3Bucket items from a CloudFormation template
func GetAllAWSS3BucketResources(template *Template) map[string]*AWSS3Bucket {

	results := map[string]*AWSS3Bucket{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3BucketWithName retrieves all AWSS3Bucket items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSS3BucketWithName(name string, template *Template) (*AWSS3Bucket, error) {

	result := &AWSS3Bucket{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket{}, errors.New("resource not found")

}
