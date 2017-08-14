package resources

// AWS::WAFRegional::XssMatchSet.XssMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html
type AWSWAFRegionalXssMatchSet_XssMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-fieldtomatch

	FieldToMatch AWSWAFRegionalXssMatchSet_FieldToMatch `json:"FieldToMatch"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::XssMatchSet.XssMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
