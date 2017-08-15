package resources

// AWS::WAF::Rule.Predicate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html
type AWSWAFRule_Predicate struct {

	// DataId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-dataid
	DataId string `json:"DataId"`

	// Negated AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-negated
	Negated bool `json:"Negated"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRule_Predicate) AWSCloudFormationType() string {
	return "AWS::WAF::Rule.Predicate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRule_Predicate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
