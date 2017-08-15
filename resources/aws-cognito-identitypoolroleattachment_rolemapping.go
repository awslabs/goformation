package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::IdentityPoolRoleAttachment.RoleMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html
type AWSCognitoIdentityPoolRoleAttachment_RoleMapping struct {

	// AmbiguousRoleResolution AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-ambiguousroleresolution
	AmbiguousRoleResolution string `json:"AmbiguousRoleResolution"`

	// RulesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-rulesconfiguration
	RulesConfiguration AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType `json:"RulesConfiguration"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment_RoleMapping) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPoolRoleAttachment.RoleMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPoolRoleAttachment_RoleMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPoolRoleAttachment_RoleMappingResources retrieves all AWSCognitoIdentityPoolRoleAttachment_RoleMapping items from a CloudFormation template
func GetAllAWSCognitoIdentityPoolRoleAttachment_RoleMapping(template *Template) map[string]*AWSCognitoIdentityPoolRoleAttachment_RoleMapping {

	results := map[string]*AWSCognitoIdentityPoolRoleAttachment_RoleMapping{}
	for name, resource := range template.Resources {
		result := &AWSCognitoIdentityPoolRoleAttachment_RoleMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPoolRoleAttachment_RoleMappingWithName retrieves all AWSCognitoIdentityPoolRoleAttachment_RoleMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoIdentityPoolRoleAttachment_RoleMapping(name string, template *Template) (*AWSCognitoIdentityPoolRoleAttachment_RoleMapping, error) {

	result := &AWSCognitoIdentityPoolRoleAttachment_RoleMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPoolRoleAttachment_RoleMapping{}, errors.New("resource not found")

}
