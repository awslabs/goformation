package resources

// AWSWAFIPSet_IPSetDescriptor AWS CloudFormation Resource (AWS::WAF::IPSet.IPSetDescriptor)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html
type AWSWAFIPSet_IPSetDescriptor struct {

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
func (r *AWSWAFIPSet_IPSetDescriptor) AWSCloudFormationType() string {
	return "AWS::WAF::IPSet.IPSetDescriptor"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFIPSet_IPSetDescriptor) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
