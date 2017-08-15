package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.WebsiteConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
type AWSS3Bucket_WebsiteConfiguration struct {

	// ErrorDocument AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-errordocument

	ErrorDocument string `json:"ErrorDocument"`

	// IndexDocument AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-indexdocument

	IndexDocument string `json:"IndexDocument"`

	// RedirectAllRequestsTo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-redirectallrequeststo

	RedirectAllRequestsTo AWSS3Bucket_RedirectAllRequestsTo `json:"RedirectAllRequestsTo"`

	// RoutingRules AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-routingrules

	RoutingRules []AWSS3Bucket_RoutingRule `json:"RoutingRules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_WebsiteConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.WebsiteConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_WebsiteConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_WebsiteConfigurationResources retrieves all AWSS3Bucket_WebsiteConfiguration items from a CloudFormation template
func GetAllAWSS3Bucket_WebsiteConfiguration(template *Template) map[string]*AWSS3Bucket_WebsiteConfiguration {

	results := map[string]*AWSS3Bucket_WebsiteConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_WebsiteConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_WebsiteConfigurationWithName retrieves all AWSS3Bucket_WebsiteConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_WebsiteConfiguration(name string, template *Template) (*AWSS3Bucket_WebsiteConfiguration, error) {

	result := &AWSS3Bucket_WebsiteConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_WebsiteConfiguration{}, errors.New("resource not found")

}
