package resources

// AWS::Cognito::UserPool.NumberAttributeConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html
type AWSCognitoUserPoolNumberAttributeConstraints struct {
    
    // MaxValue AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-maxvalue
    MaxValue string `json:"MaxValue"`
    
    // MinValue AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-minvalue
    MinValue string `json:"MinValue"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolNumberAttributeConstraints) AWSCloudFormationType() string {
    return "AWS::Cognito::UserPool.NumberAttributeConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolNumberAttributeConstraints) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}