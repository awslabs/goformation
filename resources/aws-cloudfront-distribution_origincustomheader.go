package resources

// AWS::CloudFront::Distribution.OriginCustomHeader AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin-origincustomheader.html
type AWSCloudFrontDistributionOriginCustomHeader struct {
    
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
func (r *AWSCloudFrontDistributionOriginCustomHeader) AWSCloudFormationType() string {
    return "AWS::CloudFront::Distribution.OriginCustomHeader"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistributionOriginCustomHeader) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}