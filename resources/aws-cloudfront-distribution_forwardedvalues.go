package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.ForwardedValues AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html
type AWSCloudFrontDistribution_ForwardedValues struct {

	// Cookies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-cookies

	Cookies AWSCloudFrontDistribution_Cookies `json:"Cookies"`

	// Headers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-headers

	Headers []string `json:"Headers"`

	// QueryString AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-querystring

	QueryString bool `json:"QueryString"`

	// QueryStringCacheKeys AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html#cfn-cloudfront-forwardedvalues-querystringcachekeys

	QueryStringCacheKeys []string `json:"QueryStringCacheKeys"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_ForwardedValues) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.ForwardedValues"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_ForwardedValues) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_ForwardedValuesResources retrieves all AWSCloudFrontDistribution_ForwardedValues items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_ForwardedValues(template *Template) map[string]*AWSCloudFrontDistribution_ForwardedValues {

	results := map[string]*AWSCloudFrontDistribution_ForwardedValues{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_ForwardedValues{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_ForwardedValuesWithName retrieves all AWSCloudFrontDistribution_ForwardedValues items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_ForwardedValues(name string, template *Template) (*AWSCloudFrontDistribution_ForwardedValues, error) {

	result := &AWSCloudFrontDistribution_ForwardedValues{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_ForwardedValues{}, errors.New("resource not found")

}
