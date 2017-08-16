package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCognitoUserPoolGroup AWS CloudFormation Resource (AWS::Cognito::UserPoolGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html
type AWSCognitoUserPoolGroup struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-description
	Description string `json:"Description"`

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-groupname
	GroupName string `json:"GroupName"`

	// Precedence AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-precedence
	Precedence float64 `json:"Precedence"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-rolearn
	RoleArn string `json:"RoleArn"`

	// UserPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-userpoolid
	UserPoolId string `json:"UserPoolId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolGroup) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPoolGroupResources retrieves all AWSCognitoUserPoolGroup items from a CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolGroupResources() map[string]*AWSCognitoUserPoolGroup {

	results := map[string]*AWSCognitoUserPoolGroup{}
	for name, resource := range t.Resources {
		result := &AWSCognitoUserPoolGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPoolGroupWithName retrieves all AWSCognitoUserPoolGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolGroupWithName(name string) (*AWSCognitoUserPoolGroup, error) {

	result := &AWSCognitoUserPoolGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPoolGroup{}, errors.New("resource not found")

}
