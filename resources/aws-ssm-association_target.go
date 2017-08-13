package resources

// AWS::SSM::Association.Target AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
type AWSSSMAssociationTarget struct {
    
    // Key AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key
    Key string `json:"Key"`
    
    // Values AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values
    Values []string `json:"Values"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociationTarget) AWSCloudFormationType() string {
    return "AWS::SSM::Association.Target"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMAssociationTarget) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}