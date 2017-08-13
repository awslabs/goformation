package resources

// AWS::Cognito::IdentityPoolRoleAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
type AWSCognitoIdentityPoolRoleAttachment struct {
    
    // IdentityPoolId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-identitypoolid
    IdentityPoolId string `json:"IdentityPoolId"`
    
    // RoleMappings AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-rolemappings
    RoleMappings string `json:"RoleMappings"`
    
    // Roles AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-roles
    Roles string `json:"Roles"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment) AWSCloudFormationType() string {
    return "AWS::Cognito::IdentityPoolRoleAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPoolRoleAttachment) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}