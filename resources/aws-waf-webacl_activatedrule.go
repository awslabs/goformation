package resources

// AWS::WAF::WebACL.ActivatedRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html
type AWSWAFWebACLActivatedRule struct {
    
    // Action AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-action
    Action AWSWAFWebACLActivatedRuleWafAction `json:"Action"`
    
    // Priority AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-priority
    Priority int64 `json:"Priority"`
    
    // RuleId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-ruleid
    RuleId string `json:"RuleId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACLActivatedRule) AWSCloudFormationType() string {
    return "AWS::WAF::WebACL.ActivatedRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFWebACLActivatedRule) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}