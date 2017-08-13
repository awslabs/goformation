package resources

// AWS::WAFRegional::Rule.Predicate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html
type AWSWAFRegionalRulePredicate struct {
    
    // DataId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-dataid
    DataId string `json:"DataId"`
    
    // Negated AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-negated
    Negated bool `json:"Negated"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalRulePredicate) AWSCloudFormationType() string {
    return "AWS::WAFRegional::Rule.Predicate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalRulePredicate) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}