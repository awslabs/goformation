package resources

// AWS::WAF::WebACL.WafAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html
type AWSWAFWebACLWafAction struct {
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html#cfn-waf-webacl-action-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACLWafAction) AWSCloudFormationType() string {
    return "AWS::WAF::WebACL.WafAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFWebACLWafAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}