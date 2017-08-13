package resources

// AWS::WAF::ByteMatchSet.ByteMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html
type AWSWAFByteMatchSetByteMatchTuple struct {
    
    // FieldToMatch AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch
    FieldToMatch AWSWAFByteMatchSetByteMatchTupleFieldToMatch `json:"FieldToMatch"`
    
    // PositionalConstraint AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-positionalconstraint
    PositionalConstraint string `json:"PositionalConstraint"`
    
    // TargetString AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstring
    TargetString string `json:"TargetString"`
    
    // TargetStringBase64 AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstringbase64
    TargetStringBase64 string `json:"TargetStringBase64"`
    
    // TextTransformation AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-texttransformation
    TextTransformation string `json:"TextTransformation"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFByteMatchSetByteMatchTuple) AWSCloudFormationType() string {
    return "AWS::WAF::ByteMatchSet.ByteMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFByteMatchSetByteMatchTuple) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}