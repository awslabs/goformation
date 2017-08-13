package resources

// AWS::Cognito::UserPoolUser AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
type AWSCognitoUserPoolUser struct {
    
    // DesiredDeliveryMediums AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-desireddeliverymediums
    DesiredDeliveryMediums []string `json:"DesiredDeliveryMediums"`
    
    // ForceAliasCreation AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-forcealiascreation
    ForceAliasCreation bool `json:"ForceAliasCreation"`
    
    // MessageAction AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-messageaction
    MessageAction string `json:"MessageAction"`
    
    // UserAttributes AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userattributes
    UserAttributes []AWSCognitoUserPoolUserAttributeType `json:"UserAttributes"`
    
    // UserPoolId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userpoolid
    UserPoolId string `json:"UserPoolId"`
    
    // Username AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-username
    Username string `json:"Username"`
    
    // ValidationData AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-validationdata
    ValidationData []AWSCognitoUserPoolUserAttributeType `json:"ValidationData"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUser) AWSCloudFormationType() string {
    return "AWS::Cognito::UserPoolUser"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolUser) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}