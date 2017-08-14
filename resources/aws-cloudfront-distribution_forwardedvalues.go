package resources

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
