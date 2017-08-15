package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.RedirectAllRequestsTo AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html
type AWSS3Bucket_RedirectAllRequestsTo struct {

	// HostName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html#cfn-s3-websiteconfiguration-redirectallrequeststo-hostname
	HostName string `json:"HostName"`

	// Protocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html#cfn-s3-websiteconfiguration-redirectallrequeststo-protocol
	Protocol string `json:"Protocol"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RedirectAllRequestsTo) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RedirectAllRequestsTo"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RedirectAllRequestsTo) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_RedirectAllRequestsToResources retrieves all AWSS3Bucket_RedirectAllRequestsTo items from a CloudFormation template
func GetAllAWSS3Bucket_RedirectAllRequestsTo(template *Template) map[string]*AWSS3Bucket_RedirectAllRequestsTo {

	results := map[string]*AWSS3Bucket_RedirectAllRequestsTo{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_RedirectAllRequestsTo{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_RedirectAllRequestsToWithName retrieves all AWSS3Bucket_RedirectAllRequestsTo items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_RedirectAllRequestsTo(name string, template *Template) (*AWSS3Bucket_RedirectAllRequestsTo, error) {

	result := &AWSS3Bucket_RedirectAllRequestsTo{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_RedirectAllRequestsTo{}, errors.New("resource not found")

}
