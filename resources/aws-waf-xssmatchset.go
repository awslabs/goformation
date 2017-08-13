package resources

// AWS::WAF::XssMatchSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html
type AWSWAFXssMatchSet struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-name
    Name string `json:"Name"`
    
    // XssMatchTuples AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-xssmatchtuples
    XssMatchTuples []AWSWAFXssMatchSetXssMatchTuple `json:"XssMatchTuples"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet) AWSCloudFormationType() string {
    return "AWS::WAF::XssMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFXssMatchSet) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}