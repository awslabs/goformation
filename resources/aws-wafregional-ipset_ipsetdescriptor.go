package resources

// AWS::WAFRegional::IPSet.IPSetDescriptor AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html
type AWSWAFRegionalIPSetIPSetDescriptor struct {
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html#cfn-wafregional-ipset-ipsetdescriptor-type
    Type string `json:"Type"`
    
    // Value AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html#cfn-wafregional-ipset-ipsetdescriptor-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalIPSetIPSetDescriptor) AWSCloudFormationType() string {
    return "AWS::WAFRegional::IPSet.IPSetDescriptor"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalIPSetIPSetDescriptor) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}