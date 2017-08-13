package resources

// AWS::CloudFront::Distribution.Logging AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html
type AWSCloudFrontDistributionLogging struct {
    
    // Bucket AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-bucket
    Bucket string `json:"Bucket"`
    
    // IncludeCookies AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-includecookies
    IncludeCookies bool `json:"IncludeCookies"`
    
    // Prefix AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html#cfn-cloudfront-logging-prefix
    Prefix string `json:"Prefix"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistributionLogging) AWSCloudFormationType() string {
    return "AWS::CloudFront::Distribution.Logging"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistributionLogging) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}