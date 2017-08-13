package resources

// AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html
type AWSCognitoIdentityPoolRoleAttachmentRulesConfigurationType struct {
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachmentRulesConfigurationType) AWSCloudFormationType() string {
    return "AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPoolRoleAttachmentRulesConfigurationType) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}