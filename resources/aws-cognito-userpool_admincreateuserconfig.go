package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.AdminCreateUserConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html
type AWSCognitoUserPool_AdminCreateUserConfig struct {

	// AllowAdminCreateUserOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-allowadmincreateuseronly
	AllowAdminCreateUserOnly bool `json:"AllowAdminCreateUserOnly"`

	// InviteMessageTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-invitemessagetemplate
	InviteMessageTemplate AWSCognitoUserPool_InviteMessageTemplate `json:"InviteMessageTemplate"`

	// UnusedAccountValidityDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-unusedaccountvaliditydays
	UnusedAccountValidityDays float64 `json:"UnusedAccountValidityDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_AdminCreateUserConfig) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.AdminCreateUserConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_AdminCreateUserConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_AdminCreateUserConfigResources retrieves all AWSCognitoUserPool_AdminCreateUserConfig items from a CloudFormation template
func GetAllAWSCognitoUserPool_AdminCreateUserConfig(template *Template) map[string]*AWSCognitoUserPool_AdminCreateUserConfig {

	results := map[string]*AWSCognitoUserPool_AdminCreateUserConfig{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_AdminCreateUserConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_AdminCreateUserConfigWithName retrieves all AWSCognitoUserPool_AdminCreateUserConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_AdminCreateUserConfig(name string, template *Template) (*AWSCognitoUserPool_AdminCreateUserConfig, error) {

	result := &AWSCognitoUserPool_AdminCreateUserConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_AdminCreateUserConfig{}, errors.New("resource not found")

}
