package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.Cookies AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html
type AWSCloudFrontDistribution_Cookies struct {

	// Forward AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html#cfn-cloudfront-forwardedvalues-cookies-forward
	Forward string `json:"Forward"`

	// WhitelistedNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html#cfn-cloudfront-forwardedvalues-cookies-whitelistednames
	WhitelistedNames []string `json:"WhitelistedNames"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_Cookies) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.Cookies"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_Cookies) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_CookiesResources retrieves all AWSCloudFrontDistribution_Cookies items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_Cookies(template *Template) map[string]*AWSCloudFrontDistribution_Cookies {

	results := map[string]*AWSCloudFrontDistribution_Cookies{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_Cookies{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_CookiesWithName retrieves all AWSCloudFrontDistribution_Cookies items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_Cookies(name string, template *Template) (*AWSCloudFrontDistribution_Cookies, error) {

	result := &AWSCloudFrontDistribution_Cookies{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_Cookies{}, errors.New("resource not found")

}
