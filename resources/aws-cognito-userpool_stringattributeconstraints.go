package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.StringAttributeConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
type AWSCognitoUserPool_StringAttributeConstraints struct {

	// MaxLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-maxlength

	MaxLength string `json:"MaxLength"`

	// MinLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-minlength

	MinLength string `json:"MinLength"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_StringAttributeConstraints) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.StringAttributeConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_StringAttributeConstraints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_StringAttributeConstraintsResources retrieves all AWSCognitoUserPool_StringAttributeConstraints items from a CloudFormation template
func GetAllAWSCognitoUserPool_StringAttributeConstraints(template *Template) map[string]*AWSCognitoUserPool_StringAttributeConstraints {

	results := map[string]*AWSCognitoUserPool_StringAttributeConstraints{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_StringAttributeConstraints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_StringAttributeConstraintsWithName retrieves all AWSCognitoUserPool_StringAttributeConstraints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_StringAttributeConstraints(name string, template *Template) (*AWSCognitoUserPool_StringAttributeConstraints, error) {

	result := &AWSCognitoUserPool_StringAttributeConstraints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_StringAttributeConstraints{}, errors.New("resource not found")

}
