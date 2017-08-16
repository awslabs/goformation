package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSServerlessFunction AWS CloudFormation Resource (AWS::Serverless::Function)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
type AWSServerlessFunction struct {

	// CodeUri AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	CodeUri AWSServerlessFunction_StringOrS3Location `json:"CodeUri"`

	// Handler AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Handler string `json:"Handler"`

	// Runtime AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Runtime string `json:"Runtime"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction) AWSCloudFormationType() string {
	return "AWS::Serverless::Function"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSServerlessFunction) AWSCloudFormationSpecificationVersion() string {
	return "2016-10-31"
}

// GetAllAWSServerlessFunctionResources retrieves all AWSServerlessFunction items from a CloudFormation template
func (t *Template) GetAllAWSServerlessFunctionResources() map[string]*AWSServerlessFunction {

	results := map[string]*AWSServerlessFunction{}
	for name, resource := range t.Resources {
		result := &AWSServerlessFunction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSServerlessFunctionWithName retrieves all AWSServerlessFunction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessFunctionWithName(name string) (*AWSServerlessFunction, error) {

	result := &AWSServerlessFunction{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSServerlessFunction{}, errors.New("resource not found")

}
