package resources

// AWS::SSM::Association.ParameterValues AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html
type AWSSSMAssociationParameterValues struct {
    
    // ParameterValues AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
    ParameterValues []string `json:"ParameterValues"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociationParameterValues) AWSCloudFormationType() string {
    return "AWS::SSM::Association.ParameterValues"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMAssociationParameterValues) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}