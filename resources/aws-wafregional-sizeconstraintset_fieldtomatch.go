package resources

// AWS::WAFRegional::SizeConstraintSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html
type AWSWAFRegionalSizeConstraintSetFieldToMatch struct {
    
    // Data AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html#cfn-wafregional-sizeconstraintset-fieldtomatch-data
    Data string `json:"Data"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html#cfn-wafregional-sizeconstraintset-fieldtomatch-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSizeConstraintSetFieldToMatch) AWSCloudFormationType() string {
    return "AWS::WAFRegional::SizeConstraintSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSizeConstraintSetFieldToMatch) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}