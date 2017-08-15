package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Lambda::Function.DeadLetterConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html
type AWSLambdaFunction_DeadLetterConfig struct {

	// TargetArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html#cfn-lambda-function-deadletterconfig-targetarn
	TargetArn string `json:"TargetArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_DeadLetterConfig) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.DeadLetterConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunction_DeadLetterConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaFunction_DeadLetterConfigResources retrieves all AWSLambdaFunction_DeadLetterConfig items from a CloudFormation template
func GetAllAWSLambdaFunction_DeadLetterConfig(template *Template) map[string]*AWSLambdaFunction_DeadLetterConfig {

	results := map[string]*AWSLambdaFunction_DeadLetterConfig{}
	for name, resource := range template.Resources {
		result := &AWSLambdaFunction_DeadLetterConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaFunction_DeadLetterConfigWithName retrieves all AWSLambdaFunction_DeadLetterConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLambdaFunction_DeadLetterConfig(name string, template *Template) (*AWSLambdaFunction_DeadLetterConfig, error) {

	result := &AWSLambdaFunction_DeadLetterConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaFunction_DeadLetterConfig{}, errors.New("resource not found")

}
