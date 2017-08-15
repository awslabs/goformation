package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::S3::Bucket.RedirectRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html
type AWSS3Bucket_RedirectRule struct {

	// HostName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-hostname

	HostName string `json:"HostName"`

	// HttpRedirectCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-httpredirectcode

	HttpRedirectCode string `json:"HttpRedirectCode"`

	// Protocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-protocol

	Protocol string `json:"Protocol"`

	// ReplaceKeyPrefixWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeyprefixwith

	ReplaceKeyPrefixWith string `json:"ReplaceKeyPrefixWith"`

	// ReplaceKeyWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeywith

	ReplaceKeyWith string `json:"ReplaceKeyWith"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RedirectRule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RedirectRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RedirectRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSS3Bucket_RedirectRuleResources retrieves all AWSS3Bucket_RedirectRule items from a CloudFormation template
func GetAllAWSS3Bucket_RedirectRule(template *Template) map[string]*AWSS3Bucket_RedirectRule {

	results := map[string]*AWSS3Bucket_RedirectRule{}
	for name, resource := range template.Resources {
		result := &AWSS3Bucket_RedirectRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSS3Bucket_RedirectRuleWithName retrieves all AWSS3Bucket_RedirectRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSS3Bucket_RedirectRule(name string, template *Template) (*AWSS3Bucket_RedirectRule, error) {

	result := &AWSS3Bucket_RedirectRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSS3Bucket_RedirectRule{}, errors.New("resource not found")

}
