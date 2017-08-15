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
