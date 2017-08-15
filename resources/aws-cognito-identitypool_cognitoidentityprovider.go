package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::IdentityPool.CognitoIdentityProvider AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html
type AWSCognitoIdentityPool_CognitoIdentityProvider struct {

	// ClientId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-clientid
	ClientId string `json:"ClientId"`

	// ProviderName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-providername
	ProviderName string `json:"ProviderName"`

	// ServerSideTokenCheck AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-serversidetokencheck
	ServerSideTokenCheck bool `json:"ServerSideTokenCheck"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPool_CognitoIdentityProvider) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPool.CognitoIdentityProvider"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPool_CognitoIdentityProvider) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPool_CognitoIdentityProviderResources retrieves all AWSCognitoIdentityPool_CognitoIdentityProvider items from a CloudFormation template
func GetAllAWSCognitoIdentityPool_CognitoIdentityProvider(template *Template) map[string]*AWSCognitoIdentityPool_CognitoIdentityProvider {

	results := map[string]*AWSCognitoIdentityPool_CognitoIdentityProvider{}
	for name, resource := range template.Resources {
		result := &AWSCognitoIdentityPool_CognitoIdentityProvider{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPool_CognitoIdentityProviderWithName retrieves all AWSCognitoIdentityPool_CognitoIdentityProvider items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoIdentityPool_CognitoIdentityProvider(name string, template *Template) (*AWSCognitoIdentityPool_CognitoIdentityProvider, error) {

	result := &AWSCognitoIdentityPool_CognitoIdentityProvider{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPool_CognitoIdentityProvider{}, errors.New("resource not found")

}
