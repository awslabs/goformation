package resources

// AWS::Cognito::IdentityPool AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
type AWSCognitoIdentityPool struct {

	// AllowUnauthenticatedIdentities AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	AllowUnauthenticatedIdentities bool `json:"AllowUnauthenticatedIdentities"`

	// CognitoEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	CognitoEvents object `json:"CognitoEvents"`

	// CognitoIdentityProviders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	CognitoIdentityProviders []AWSCognitoIdentityPoolCognitoIdentityProvider `json:"CognitoIdentityProviders"`

	// CognitoStreams AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	CognitoStreams AWSCognitoIdentityPoolCognitoStreams `json:"CognitoStreams"`

	// DeveloperProviderName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	DeveloperProviderName string `json:"DeveloperProviderName"`

	// IdentityPoolName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	IdentityPoolName string `json:"IdentityPoolName"`

	// OpenIdConnectProviderARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	OpenIdConnectProviderARNs []AWSCognitoIdentityPoolstring `json:"OpenIdConnectProviderARNs"`

	// PushSync AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	PushSync AWSCognitoIdentityPoolPushSync `json:"PushSync"`

	// SamlProviderARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	SamlProviderARNs []AWSCognitoIdentityPoolstring `json:"SamlProviderARNs"`

	// SupportedLoginProviders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	SupportedLoginProviders object `json:"SupportedLoginProviders"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPool) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPool"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPool) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
