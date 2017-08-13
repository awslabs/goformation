package resources

// AWS::Cognito::UserPool.Policies AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
type AWSCognitoUserPoolPolicies struct {
    
    // PasswordPolicy AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
    PasswordPolicy AWSCognitoUserPoolPoliciesPasswordPolicy `json:"PasswordPolicy"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolPolicies) AWSCloudFormationType() string {
    return "AWS::Cognito::UserPool.Policies"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolPolicies) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}