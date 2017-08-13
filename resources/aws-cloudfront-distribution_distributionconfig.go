package resources

// AWS::CloudFront::Distribution.DistributionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html
type AWSCloudFrontDistributionDistributionConfig struct {
    
    // Aliases AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-aliases
    Aliases []string `json:"Aliases"`
    
    // CacheBehaviors AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-cachebehaviors
    CacheBehaviors []AWSCloudFrontDistributionDistributionConfigCacheBehavior `json:"CacheBehaviors"`
    
    // Comment AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-comment
    Comment string `json:"Comment"`
    
    // CustomErrorResponses AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-customerrorresponses
    CustomErrorResponses []AWSCloudFrontDistributionDistributionConfigCustomErrorResponse `json:"CustomErrorResponses"`
    
    // DefaultCacheBehavior AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-defaultcachebehavior
    DefaultCacheBehavior AWSCloudFrontDistributionDistributionConfigDefaultCacheBehavior `json:"DefaultCacheBehavior"`
    
    // DefaultRootObject AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-defaultrootobject
    DefaultRootObject string `json:"DefaultRootObject"`
    
    // Enabled AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-enabled
    Enabled bool `json:"Enabled"`
    
    // HttpVersion AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-httpversion
    HttpVersion string `json:"HttpVersion"`
    
    // Logging AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-logging
    Logging AWSCloudFrontDistributionDistributionConfigLogging `json:"Logging"`
    
    // Origins AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-origins
    Origins []AWSCloudFrontDistributionDistributionConfigOrigin `json:"Origins"`
    
    // PriceClass AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-priceclass
    PriceClass string `json:"PriceClass"`
    
    // Restrictions AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-restrictions
    Restrictions AWSCloudFrontDistributionDistributionConfigRestrictions `json:"Restrictions"`
    
    // ViewerCertificate AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-viewercertificate
    ViewerCertificate AWSCloudFrontDistributionDistributionConfigViewerCertificate `json:"ViewerCertificate"`
    
    // WebACLId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html#cfn-cloudfront-distributionconfig-webaclid
    WebACLId string `json:"WebACLId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistributionDistributionConfig) AWSCloudFormationType() string {
    return "AWS::CloudFront::Distribution.DistributionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistributionDistributionConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}