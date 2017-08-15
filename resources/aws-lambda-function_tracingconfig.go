package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Lambda::Function.TracingConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html
type AWSLambdaFunction_TracingConfig struct {

	// Mode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tracingconfig.html#cfn-lambda-function-tracingconfig-mode

	Mode string `json:"Mode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_TracingConfig) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.TracingConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunction_TracingConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaFunction_TracingConfigResources retrieves all AWSLambdaFunction_TracingConfig items from a CloudFormation template
func GetAllAWSLambdaFunction_TracingConfig(template *Template) map[string]*AWSLambdaFunction_TracingConfig {

	results := map[string]*AWSLambdaFunction_TracingConfig{}
	for name, resource := range template.Resources {
		result := &AWSLambdaFunction_TracingConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaFunction_TracingConfigWithName retrieves all AWSLambdaFunction_TracingConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLambdaFunction_TracingConfig(name string, template *Template) (*AWSLambdaFunction_TracingConfig, error) {

	result := &AWSLambdaFunction_TracingConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaFunction_TracingConfig{}, errors.New("resource not found")

}
