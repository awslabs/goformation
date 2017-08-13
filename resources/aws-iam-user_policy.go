package resources

// AWS::IAM::User.Policy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type AWSIAMUserPolicy struct {
    
    // PolicyDocument AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
    PolicyDocument string `json:"PolicyDocument"`
    
    // PolicyName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
    PolicyName string `json:"PolicyName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUserPolicy) AWSCloudFormationType() string {
    return "AWS::IAM::User.Policy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUserPolicy) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}