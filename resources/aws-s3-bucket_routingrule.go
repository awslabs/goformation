package resources

// AWS::S3::Bucket.RoutingRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
type AWSS3BucketRoutingRule struct {
    
    // RedirectRule AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-redirectrule
    RedirectRule AWSS3BucketRoutingRuleRedirectRule `json:"RedirectRule"`
    
    // RoutingRuleCondition AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition
    RoutingRuleCondition AWSS3BucketRoutingRuleRoutingRuleCondition `json:"RoutingRuleCondition"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketRoutingRule) AWSCloudFormationType() string {
    return "AWS::S3::Bucket.RoutingRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketRoutingRule) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}