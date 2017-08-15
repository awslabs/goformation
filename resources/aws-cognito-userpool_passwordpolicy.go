package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.PasswordPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html
type AWSCognitoUserPool_PasswordPolicy struct {

	// MinimumLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-minimumlength

	MinimumLength int64 `json:"MinimumLength"`

	// RequireLowercase AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requirelowercase

	RequireLowercase bool `json:"RequireLowercase"`

	// RequireNumbers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requirenumbers

	RequireNumbers bool `json:"RequireNumbers"`

	// RequireSymbols AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requiresymbols

	RequireSymbols bool `json:"RequireSymbols"`

	// RequireUppercase AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-passwordpolicy.html#cfn-cognito-userpool-passwordpolicy-requireuppercase

	RequireUppercase bool `json:"RequireUppercase"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_PasswordPolicy) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.PasswordPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_PasswordPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_PasswordPolicyResources retrieves all AWSCognitoUserPool_PasswordPolicy items from a CloudFormation template
func GetAllAWSCognitoUserPool_PasswordPolicy(template *Template) map[string]*AWSCognitoUserPool_PasswordPolicy {

	results := map[string]*AWSCognitoUserPool_PasswordPolicy{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_PasswordPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_PasswordPolicyWithName retrieves all AWSCognitoUserPool_PasswordPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_PasswordPolicy(name string, template *Template) (*AWSCognitoUserPool_PasswordPolicy, error) {

	result := &AWSCognitoUserPool_PasswordPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_PasswordPolicy{}, errors.New("resource not found")

}
