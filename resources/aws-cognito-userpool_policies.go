package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.Policies AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
type AWSCognitoUserPool_Policies struct {

	// PasswordPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
	PasswordPolicy AWSCognitoUserPool_PasswordPolicy `json:"PasswordPolicy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_Policies) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.Policies"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_Policies) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_PoliciesResources retrieves all AWSCognitoUserPool_Policies items from a CloudFormation template
func GetAllAWSCognitoUserPool_Policies(template *Template) map[string]*AWSCognitoUserPool_Policies {

	results := map[string]*AWSCognitoUserPool_Policies{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_Policies{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_PoliciesWithName retrieves all AWSCognitoUserPool_Policies items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_Policies(name string, template *Template) (*AWSCognitoUserPool_Policies, error) {

	result := &AWSCognitoUserPool_Policies{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_Policies{}, errors.New("resource not found")

}
