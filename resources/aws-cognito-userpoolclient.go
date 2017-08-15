package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCognitoUserPoolClient AWS CloudFormation Resource (AWS::Cognito::UserPoolClient)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
type AWSCognitoUserPoolClient struct {

	// ClientName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-clientname
	ClientName string `json:"ClientName"`

	// ExplicitAuthFlows AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-explicitauthflows
	ExplicitAuthFlows []string `json:"ExplicitAuthFlows"`

	// GenerateSecret AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-generatesecret
	GenerateSecret bool `json:"GenerateSecret"`

	// ReadAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-readattributes
	ReadAttributes []string `json:"ReadAttributes"`

	// RefreshTokenValidity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-refreshtokenvalidity
	RefreshTokenValidity float64 `json:"RefreshTokenValidity"`

	// UserPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-userpoolid
	UserPoolId string `json:"UserPoolId"`

	// WriteAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-writeattributes
	WriteAttributes []string `json:"WriteAttributes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolClient) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolClient"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolClient) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPoolClientResources retrieves all AWSCognitoUserPoolClient items from a CloudFormation template
func GetAllAWSCognitoUserPoolClientResources(template *Template) map[string]*AWSCognitoUserPoolClient {

	results := map[string]*AWSCognitoUserPoolClient{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPoolClient{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPoolClientWithName retrieves all AWSCognitoUserPoolClient items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCognitoUserPoolClientWithName(name string, template *Template) (*AWSCognitoUserPoolClient, error) {

	result := &AWSCognitoUserPoolClient{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPoolClient{}, errors.New("resource not found")

}
