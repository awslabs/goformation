package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Lambda::Alias AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html
type AWSLambdaAlias struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-description
	Description string `json:"Description"`

	// FunctionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionname
	FunctionName string `json:"FunctionName"`

	// FunctionVersion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-functionversion
	FunctionVersion string `json:"FunctionVersion"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html#cfn-lambda-alias-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaAlias) AWSCloudFormationType() string {
	return "AWS::Lambda::Alias"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaAlias) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaAliasResources retrieves all AWSLambdaAlias items from a CloudFormation template
func GetAllAWSLambdaAliasResources(template *Template) map[string]*AWSLambdaAlias {

	results := map[string]*AWSLambdaAlias{}
	for name, resource := range template.Resources {
		result := &AWSLambdaAlias{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaAliasWithName retrieves all AWSLambdaAlias items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSLambdaAliasWithName(name string, template *Template) (*AWSLambdaAlias, error) {

	result := &AWSLambdaAlias{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaAlias{}, errors.New("resource not found")

}
