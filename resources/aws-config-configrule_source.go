package resources

// AWS::Config::ConfigRule.Source AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html
type AWSConfigConfigRuleSource struct {
    
    // Owner AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-owner
    Owner string `json:"Owner"`
    
    // SourceDetails AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourcedetails
    SourceDetails []AWSConfigConfigRuleSourceSourceDetail `json:"SourceDetails"`
    
    // SourceIdentifier AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourceidentifier
    SourceIdentifier string `json:"SourceIdentifier"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRuleSource) AWSCloudFormationType() string {
    return "AWS::Config::ConfigRule.Source"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRuleSource) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}