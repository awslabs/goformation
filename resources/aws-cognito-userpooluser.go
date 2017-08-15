package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPoolUser AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
type AWSCognitoUserPoolUser struct {

	// DesiredDeliveryMediums AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-desireddeliverymediums
	DesiredDeliveryMediums []string `json:"DesiredDeliveryMediums"`

	// ForceAliasCreation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-forcealiascreation
	ForceAliasCreation bool `json:"ForceAliasCreation"`

	// MessageAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-messageaction
	MessageAction string `json:"MessageAction"`

	// UserAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userattributes
	UserAttributes []AWSCognitoUserPoolUser_AttributeType `json:"UserAttributes"`

	// UserPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userpoolid
	UserPoolId string `json:"UserPoolId"`

	// Username AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-username
	Username string `json:"Username"`

	// ValidationData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-validationdata
	ValidationData []AWSCognitoUserPoolUser_AttributeType `json:"ValidationData"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUser) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolUser"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolUser) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPoolUserResources retrieves all AWSCognitoUserPoolUser items from a CloudFormation template
func GetAllAWSCognitoUserPoolUser(template *Template) map[string]*AWSCognitoUserPoolUser {

	results := map[string]*AWSCognitoUserPoolUser{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPoolUser{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPoolUserWithName retrieves all AWSCognitoUserPoolUser items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPoolUser(name string, template *Template) (*AWSCognitoUserPoolUser, error) {

	result := &AWSCognitoUserPoolUser{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPoolUser{}, errors.New("resource not found")

}
