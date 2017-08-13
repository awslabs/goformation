package resources

// AWS::WAF::IPSet.IPSetDescriptor AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html
type AWSWAFIPSetIPSetDescriptor struct {
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html#cfn-waf-ipset-ipsetdescriptors-type
    Type string `json:"Type"`
    
    // Value AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html#cfn-waf-ipset-ipsetdescriptors-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFIPSetIPSetDescriptor) AWSCloudFormationType() string {
    return "AWS::WAF::IPSet.IPSetDescriptor"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFIPSetIPSetDescriptor) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}