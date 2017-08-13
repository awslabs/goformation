package resources

// AWS::Cognito::UserPool.DeviceConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html
type AWSCognitoUserPoolDeviceConfiguration struct {
    
    // ChallengeRequiredOnNewDevice AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-challengerequiredonnewdevice
    ChallengeRequiredOnNewDevice bool `json:"ChallengeRequiredOnNewDevice"`
    
    // DeviceOnlyRememberedOnUserPrompt AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-deviceonlyrememberedonuserprompt
    DeviceOnlyRememberedOnUserPrompt bool `json:"DeviceOnlyRememberedOnUserPrompt"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolDeviceConfiguration) AWSCloudFormationType() string {
    return "AWS::Cognito::UserPool.DeviceConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolDeviceConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}