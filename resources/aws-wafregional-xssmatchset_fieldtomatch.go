package resources

// AWS::WAFRegional::XssMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html
type AWSWAFRegionalXssMatchSetFieldToMatch struct {
    
    // Data AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-data
    Data string `json:"Data"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSetFieldToMatch) AWSCloudFormationType() string {
    return "AWS::WAFRegional::XssMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalXssMatchSetFieldToMatch) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}