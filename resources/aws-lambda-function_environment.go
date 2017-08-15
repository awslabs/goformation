package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Lambda::Function.Environment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
type AWSLambdaFunction_Environment struct {

	// Variables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables

	Variables map[string]string `json:"Variables"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_Environment) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunction_Environment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaFunction_EnvironmentResources retrieves all AWSLambdaFunction_Environment items from a CloudFormation template
func GetAllAWSLambdaFunction_Environment(template *Template) map[string]*AWSLambdaFunction_Environment {

	results := map[string]*AWSLambdaFunction_Environment{}
	for name, resource := range template.Resources {
		result := &AWSLambdaFunction_Environment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaFunction_EnvironmentWithName retrieves all AWSLambdaFunction_Environment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLambdaFunction_Environment(name string, template *Template) (*AWSLambdaFunction_Environment, error) {

	result := &AWSLambdaFunction_Environment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaFunction_Environment{}, errors.New("resource not found")

}
