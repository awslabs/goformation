package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html
type AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType struct {
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPoolRoleAttachment_RulesConfigurationTypeResources retrieves all AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType items from a CloudFormation template
func GetAllAWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType(template *Template) map[string]*AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType {

	results := map[string]*AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType{}
	for name, resource := range template.Resources {
		result := &AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPoolRoleAttachment_RulesConfigurationTypeWithName retrieves all AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType(name string, template *Template) (*AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType, error) {

	result := &AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType{}, errors.New("resource not found")

}
