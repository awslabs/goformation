package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.EmailConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html
type AWSCognitoUserPool_EmailConfiguration struct {

	// ReplyToEmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-replytoemailaddress
	ReplyToEmailAddress string `json:"ReplyToEmailAddress"`

	// SourceArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-sourcearn
	SourceArn string `json:"SourceArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_EmailConfiguration) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.EmailConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_EmailConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_EmailConfigurationResources retrieves all AWSCognitoUserPool_EmailConfiguration items from a CloudFormation template
func GetAllAWSCognitoUserPool_EmailConfiguration(template *Template) map[string]*AWSCognitoUserPool_EmailConfiguration {

	results := map[string]*AWSCognitoUserPool_EmailConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_EmailConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_EmailConfigurationWithName retrieves all AWSCognitoUserPool_EmailConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_EmailConfiguration(name string, template *Template) (*AWSCognitoUserPool_EmailConfiguration, error) {

	result := &AWSCognitoUserPool_EmailConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_EmailConfiguration{}, errors.New("resource not found")

}
