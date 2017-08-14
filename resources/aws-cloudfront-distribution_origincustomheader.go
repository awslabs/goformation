package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.OriginCustomHeader AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html
type AWSCloudFrontDistribution_OriginCustomHeader struct {

	// HeaderName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html#cfn-cloudfront-origin-origincustomheader-headername

	HeaderName string `json:"HeaderName"`

	// HeaderValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html#cfn-cloudfront-origin-origincustomheader-headervalue

	HeaderValue string `json:"HeaderValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_OriginCustomHeader) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.OriginCustomHeader"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_OriginCustomHeader) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_OriginCustomHeaderResources retrieves all AWSCloudFrontDistribution_OriginCustomHeader items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_OriginCustomHeader(template *Template) map[string]*AWSCloudFrontDistribution_OriginCustomHeader {

	results := map[string]*AWSCloudFrontDistribution_OriginCustomHeader{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_OriginCustomHeader{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_OriginCustomHeaderWithName retrieves all AWSCloudFrontDistribution_OriginCustomHeader items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_OriginCustomHeader(name string, template *Template) (*AWSCloudFrontDistribution_OriginCustomHeader, error) {

	result := &AWSCloudFrontDistribution_OriginCustomHeader{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_OriginCustomHeader{}, errors.New("resource not found")

}
