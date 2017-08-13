package resources

// AWS::WAFRegional::SizeConstraintSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
type AWSWAFRegionalSizeConstraintSet struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-name
    Name string `json:"Name"`
    
    // SizeConstraints AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-sizeconstraints
    SizeConstraints []AWSWAFRegionalSizeConstraintSetSizeConstraint `json:"SizeConstraints"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSizeConstraintSet) AWSCloudFormationType() string {
    return "AWS::WAFRegional::SizeConstraintSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSizeConstraintSet) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}