package resources

// AWS::WAF::WebACL AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html
type AWSWAFWebACL struct {
    
    // DefaultAction AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-defaultaction
    DefaultAction AWSWAFWebACLWafAction `json:"DefaultAction"`
    
    // MetricName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-metricname
    MetricName string `json:"MetricName"`
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-name
    Name string `json:"Name"`
    
    // Rules AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-rules
    Rules []AWSWAFWebACLActivatedRule `json:"Rules"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACL) AWSCloudFormationType() string {
    return "AWS::WAF::WebACL"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFWebACL) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}