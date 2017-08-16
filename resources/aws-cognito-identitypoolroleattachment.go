package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCognitoIdentityPoolRoleAttachment AWS CloudFormation Resource (AWS::Cognito::IdentityPoolRoleAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
type AWSCognitoIdentityPoolRoleAttachment struct {

	// IdentityPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-identitypoolid
	IdentityPoolId string `json:"IdentityPoolId"`

	// RoleMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-rolemappings
	RoleMappings interface{} `json:"RoleMappings"`

	// Roles AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-roles
	Roles interface{} `json:"Roles"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPoolRoleAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPoolRoleAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPoolRoleAttachmentResources retrieves all AWSCognitoIdentityPoolRoleAttachment items from a CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolRoleAttachmentResources() map[string]*AWSCognitoIdentityPoolRoleAttachment {

	results := map[string]*AWSCognitoIdentityPoolRoleAttachment{}
	for name, resource := range t.Resources {
		result := &AWSCognitoIdentityPoolRoleAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPoolRoleAttachmentWithName retrieves all AWSCognitoIdentityPoolRoleAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolRoleAttachmentWithName(name string) (*AWSCognitoIdentityPoolRoleAttachment, error) {

	result := &AWSCognitoIdentityPoolRoleAttachment{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPoolRoleAttachment{}, errors.New("resource not found")

}
