package resources

// AWS::Cognito::UserPool AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
type AWSCognitoUserPool struct {

	// AdminCreateUserConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-admincreateuserconfig

	AdminCreateUserConfig AWSCognitoUserPool_AdminCreateUserConfig `json:"AdminCreateUserConfig"`

	// AliasAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-aliasattributes

	AliasAttributes []string `json:"AliasAttributes"`

	// AutoVerifiedAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-autoverifiedattributes

	AutoVerifiedAttributes []string `json:"AutoVerifiedAttributes"`

	// DeviceConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-deviceconfiguration

	DeviceConfiguration AWSCognitoUserPool_DeviceConfiguration `json:"DeviceConfiguration"`

	// EmailConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailconfiguration

	EmailConfiguration AWSCognitoUserPool_EmailConfiguration `json:"EmailConfiguration"`

	// EmailVerificationMessage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationmessage

	EmailVerificationMessage string `json:"EmailVerificationMessage"`

	// EmailVerificationSubject AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-emailverificationsubject

	EmailVerificationSubject string `json:"EmailVerificationSubject"`

	// LambdaConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-lambdaconfig

	LambdaConfig AWSCognitoUserPool_LambdaConfig `json:"LambdaConfig"`

	// MfaConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-mfaconfiguration

	MfaConfiguration string `json:"MfaConfiguration"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-policies

	Policies AWSCognitoUserPool_Policies `json:"Policies"`

	// Schema AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-schema

	Schema []AWSCognitoUserPool_SchemaAttribute `json:"Schema"`

	// SmsAuthenticationMessage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsauthenticationmessage

	SmsAuthenticationMessage string `json:"SmsAuthenticationMessage"`

	// SmsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsconfiguration

	SmsConfiguration AWSCognitoUserPool_SmsConfiguration `json:"SmsConfiguration"`

	// SmsVerificationMessage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-smsverificationmessage

	SmsVerificationMessage string `json:"SmsVerificationMessage"`

	// UserPoolName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpoolname

	UserPoolName string `json:"UserPoolName"`

	// UserPoolTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html#cfn-cognito-userpool-userpooltags

	UserPoolTags interface{} `json:"UserPoolTags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
