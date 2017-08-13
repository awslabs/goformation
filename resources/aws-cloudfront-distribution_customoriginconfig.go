package resources

// AWS::CloudFront::Distribution.CustomOriginConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html
type AWSCloudFrontDistributionCustomOriginConfig struct {
    
    // HTTPPort AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-httpport
    HTTPPort int64 `json:"HTTPPort"`
    
    // HTTPSPort AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-httpsport
    HTTPSPort int64 `json:"HTTPSPort"`
    
    // OriginProtocolPolicy AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-originprotocolpolicy
    OriginProtocolPolicy string `json:"OriginProtocolPolicy"`
    
    // OriginSSLProtocols AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-originsslprotocols
    OriginSSLProtocols []string `json:"OriginSSLProtocols"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistributionCustomOriginConfig) AWSCloudFormationType() string {
    return "AWS::CloudFront::Distribution.CustomOriginConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistributionCustomOriginConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}