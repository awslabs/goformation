package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.SmsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html
type AWSCognitoUserPool_SmsConfiguration struct {

	// ExternalId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html#cfn-cognito-userpool-smsconfiguration-externalid
	ExternalId string `json:"ExternalId"`

	// SnsCallerArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html#cfn-cognito-userpool-smsconfiguration-snscallerarn
	SnsCallerArn string `json:"SnsCallerArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_SmsConfiguration) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.SmsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_SmsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_SmsConfigurationResources retrieves all AWSCognitoUserPool_SmsConfiguration items from a CloudFormation template
func GetAllAWSCognitoUserPool_SmsConfiguration(template *Template) map[string]*AWSCognitoUserPool_SmsConfiguration {

	results := map[string]*AWSCognitoUserPool_SmsConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_SmsConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_SmsConfigurationWithName retrieves all AWSCognitoUserPool_SmsConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_SmsConfiguration(name string, template *Template) (*AWSCognitoUserPool_SmsConfiguration, error) {

	result := &AWSCognitoUserPool_SmsConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_SmsConfiguration{}, errors.New("resource not found")

}
