package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.InviteMessageTemplate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html
type AWSCognitoUserPool_InviteMessageTemplate struct {

	// EmailMessage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-emailmessage
	EmailMessage string `json:"EmailMessage"`

	// EmailSubject AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-emailsubject
	EmailSubject string `json:"EmailSubject"`

	// SMSMessage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-invitemessagetemplate.html#cfn-cognito-userpool-invitemessagetemplate-smsmessage
	SMSMessage string `json:"SMSMessage"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_InviteMessageTemplate) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.InviteMessageTemplate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_InviteMessageTemplate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_InviteMessageTemplateResources retrieves all AWSCognitoUserPool_InviteMessageTemplate items from a CloudFormation template
func GetAllAWSCognitoUserPool_InviteMessageTemplate(template *Template) map[string]*AWSCognitoUserPool_InviteMessageTemplate {

	results := map[string]*AWSCognitoUserPool_InviteMessageTemplate{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_InviteMessageTemplate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_InviteMessageTemplateWithName retrieves all AWSCognitoUserPool_InviteMessageTemplate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_InviteMessageTemplate(name string, template *Template) (*AWSCognitoUserPool_InviteMessageTemplate, error) {

	result := &AWSCognitoUserPool_InviteMessageTemplate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_InviteMessageTemplate{}, errors.New("resource not found")

}
