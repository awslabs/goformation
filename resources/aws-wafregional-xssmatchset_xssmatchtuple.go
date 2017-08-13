package resources

// AWS::WAFRegional::XssMatchSet.XssMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html
type AWSWAFRegionalXssMatchSetXssMatchTuple struct {
    
    // FieldToMatch AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-fieldtomatch
    FieldToMatch AWSWAFRegionalXssMatchSetXssMatchTupleFieldToMatch `json:"FieldToMatch"`
    
    // TextTransformation AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-texttransformation
    TextTransformation string `json:"TextTransformation"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSetXssMatchTuple) AWSCloudFormationType() string {
    return "AWS::WAFRegional::XssMatchSet.XssMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalXssMatchSetXssMatchTuple) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}