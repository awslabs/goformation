package resources

// AWS::WAF::IPSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html
type AWSWAFIPSet struct {

	// IPSetDescriptors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-ipsetdescriptors

	IPSetDescriptors []AWSWAFIPSet_IPSetDescriptor `json:"IPSetDescriptors"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-name

	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFIPSet) AWSCloudFormationType() string {
	return "AWS::WAF::IPSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFIPSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
