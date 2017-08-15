package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPoolUserToGroupAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html
type AWSCognitoUserPoolUserToGroupAttachment struct {

	// GroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-groupname
	GroupName string `json:"GroupName"`

	// UserPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-userpoolid
	UserPoolId string `json:"UserPoolId"`

	// Username AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html#cfn-cognito-userpoolusertogroupattachment-username
	Username string `json:"Username"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUserToGroupAttachment) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolUserToGroupAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolUserToGroupAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPoolUserToGroupAttachmentResources retrieves all AWSCognitoUserPoolUserToGroupAttachment items from a CloudFormation template
func GetAllAWSCognitoUserPoolUserToGroupAttachment(template *Template) map[string]*AWSCognitoUserPoolUserToGroupAttachment {

	results := map[string]*AWSCognitoUserPoolUserToGroupAttachment{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPoolUserToGroupAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPoolUserToGroupAttachmentWithName retrieves all AWSCognitoUserPoolUserToGroupAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPoolUserToGroupAttachment(name string, template *Template) (*AWSCognitoUserPoolUserToGroupAttachment, error) {

	result := &AWSCognitoUserPoolUserToGroupAttachment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPoolUserToGroupAttachment{}, errors.New("resource not found")

}
