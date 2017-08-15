package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPoolUser.AttributeType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html
type AWSCognitoUserPoolUser_AttributeType struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUser_AttributeType) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolUser.AttributeType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolUser_AttributeType) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPoolUser_AttributeTypeResources retrieves all AWSCognitoUserPoolUser_AttributeType items from a CloudFormation template
func GetAllAWSCognitoUserPoolUser_AttributeType(template *Template) map[string]*AWSCognitoUserPoolUser_AttributeType {

	results := map[string]*AWSCognitoUserPoolUser_AttributeType{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPoolUser_AttributeType{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPoolUser_AttributeTypeWithName retrieves all AWSCognitoUserPoolUser_AttributeType items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPoolUser_AttributeType(name string, template *Template) (*AWSCognitoUserPoolUser_AttributeType, error) {

	result := &AWSCognitoUserPoolUser_AttributeType{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPoolUser_AttributeType{}, errors.New("resource not found")

}
