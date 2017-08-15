package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.LambdaConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html
type AWSCognitoUserPool_LambdaConfig struct {

	// CreateAuthChallenge AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-createauthchallenge
	CreateAuthChallenge string `json:"CreateAuthChallenge"`

	// CustomMessage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-custommessage
	CustomMessage string `json:"CustomMessage"`

	// DefineAuthChallenge AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-defineauthchallenge
	DefineAuthChallenge string `json:"DefineAuthChallenge"`

	// PostAuthentication AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postauthentication
	PostAuthentication string `json:"PostAuthentication"`

	// PostConfirmation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postconfirmation
	PostConfirmation string `json:"PostConfirmation"`

	// PreAuthentication AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-preauthentication
	PreAuthentication string `json:"PreAuthentication"`

	// PreSignUp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-presignup
	PreSignUp string `json:"PreSignUp"`

	// VerifyAuthChallengeResponse AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-verifyauthchallengeresponse
	VerifyAuthChallengeResponse string `json:"VerifyAuthChallengeResponse"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_LambdaConfig) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.LambdaConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_LambdaConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_LambdaConfigResources retrieves all AWSCognitoUserPool_LambdaConfig items from a CloudFormation template
func GetAllAWSCognitoUserPool_LambdaConfig(template *Template) map[string]*AWSCognitoUserPool_LambdaConfig {

	results := map[string]*AWSCognitoUserPool_LambdaConfig{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_LambdaConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_LambdaConfigWithName retrieves all AWSCognitoUserPool_LambdaConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_LambdaConfig(name string, template *Template) (*AWSCognitoUserPool_LambdaConfig, error) {

	result := &AWSCognitoUserPool_LambdaConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_LambdaConfig{}, errors.New("resource not found")

}
