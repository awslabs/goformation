package resources

// AWS::WAF::SizeConstraintSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html
type AWSWAFSizeConstraintSet struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-name
    Name string `json:"Name"`
    
    // SizeConstraints AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-sizeconstraints
    SizeConstraints []AWSWAFSizeConstraintSetSizeConstraint `json:"SizeConstraints"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSizeConstraintSet) AWSCloudFormationType() string {
    return "AWS::WAF::SizeConstraintSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSizeConstraintSet) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}