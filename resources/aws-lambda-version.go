package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSLambdaVersion AWS CloudFormation Resource (AWS::Lambda::Version)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html
type AWSLambdaVersion struct {

	// CodeSha256 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-codesha256
	CodeSha256 string `json:"CodeSha256"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-description
	Description string `json:"Description"`

	// FunctionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html#cfn-lambda-version-functionname
	FunctionName string `json:"FunctionName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaVersion) AWSCloudFormationType() string {
	return "AWS::Lambda::Version"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaVersion) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaVersionResources retrieves all AWSLambdaVersion items from a CloudFormation template
func (t *Template) GetAllAWSLambdaVersionResources() map[string]*AWSLambdaVersion {

	results := map[string]*AWSLambdaVersion{}
	for name, resource := range t.Resources {
		result := &AWSLambdaVersion{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaVersionWithName retrieves all AWSLambdaVersion items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLambdaVersionWithName(name string) (*AWSLambdaVersion, error) {

	result := &AWSLambdaVersion{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaVersion{}, errors.New("resource not found")

}
