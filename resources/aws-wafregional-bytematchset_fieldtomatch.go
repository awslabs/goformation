package resources

// AWS::WAFRegional::ByteMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html
type AWSWAFRegionalByteMatchSetFieldToMatch struct {
    
    // Data AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-data
    Data string `json:"Data"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalByteMatchSetFieldToMatch) AWSCloudFormationType() string {
    return "AWS::WAFRegional::ByteMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalByteMatchSetFieldToMatch) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}