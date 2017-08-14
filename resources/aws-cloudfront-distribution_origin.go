package resources

// AWS::CloudFront::Distribution.Origin AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html
type AWSCloudFrontDistributionOrigin struct {

	// CustomOriginConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-customorigin
	CustomOriginConfig AWSCloudFrontDistributionOriginCustomOriginConfig `json:"CustomOriginConfig"`

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-domainname
	DomainName string `json:"DomainName"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-id
	Id string `json:"Id"`

	// OriginCustomHeaders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-origincustomheaders
	OriginCustomHeaders []AWSCloudFrontDistributionOriginOriginCustomHeader `json:"OriginCustomHeaders"`

	// OriginPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-originpath
	OriginPath string `json:"OriginPath"`

	// S3OriginConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-s3origin
	S3OriginConfig AWSCloudFrontDistributionOriginS3OriginConfig `json:"S3OriginConfig"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistributionOrigin) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.Origin"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistributionOrigin) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
