package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::IdentityPoolRoleAttachment.MappingRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html
type AWSCognitoIdentityPoolRoleAttachment_MappingRule struct {

	// Claim AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-claim

	Claim string `json:"Claim"`

	// MatchType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-matchtype

	MatchType string `json:"MatchType"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-rolearn

	RoleARN string `json:"RoleARN"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment_MappingRule) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPoolRoleAttachment.MappingRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPoolRoleAttachment_MappingRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPoolRoleAttachment_MappingRuleResources retrieves all AWSCognitoIdentityPoolRoleAttachment_MappingRule items from a CloudFormation template
func GetAllAWSCognitoIdentityPoolRoleAttachment_MappingRule(template *Template) map[string]*AWSCognitoIdentityPoolRoleAttachment_MappingRule {

	results := map[string]*AWSCognitoIdentityPoolRoleAttachment_MappingRule{}
	for name, resource := range template.Resources {
		result := &AWSCognitoIdentityPoolRoleAttachment_MappingRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPoolRoleAttachment_MappingRuleWithName retrieves all AWSCognitoIdentityPoolRoleAttachment_MappingRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoIdentityPoolRoleAttachment_MappingRule(name string, template *Template) (*AWSCognitoIdentityPoolRoleAttachment_MappingRule, error) {

	result := &AWSCognitoIdentityPoolRoleAttachment_MappingRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPoolRoleAttachment_MappingRule{}, errors.New("resource not found")

}
