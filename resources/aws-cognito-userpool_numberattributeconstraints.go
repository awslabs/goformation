package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.NumberAttributeConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html
type AWSCognitoUserPool_NumberAttributeConstraints struct {

	// MaxValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-maxvalue
	MaxValue string `json:"MaxValue"`

	// MinValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-numberattributeconstraints.html#cfn-cognito-userpool-numberattributeconstraints-minvalue
	MinValue string `json:"MinValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_NumberAttributeConstraints) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.NumberAttributeConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_NumberAttributeConstraints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_NumberAttributeConstraintsResources retrieves all AWSCognitoUserPool_NumberAttributeConstraints items from a CloudFormation template
func GetAllAWSCognitoUserPool_NumberAttributeConstraints(template *Template) map[string]*AWSCognitoUserPool_NumberAttributeConstraints {

	results := map[string]*AWSCognitoUserPool_NumberAttributeConstraints{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_NumberAttributeConstraints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_NumberAttributeConstraintsWithName retrieves all AWSCognitoUserPool_NumberAttributeConstraints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_NumberAttributeConstraints(name string, template *Template) (*AWSCognitoUserPool_NumberAttributeConstraints, error) {

	result := &AWSCognitoUserPool_NumberAttributeConstraints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_NumberAttributeConstraints{}, errors.New("resource not found")

}
