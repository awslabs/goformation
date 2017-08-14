package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::App.EnvironmentVariable AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html
type AWSOpsWorksApp_EnvironmentVariable struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-key

	Key string `json:"Key"`

	// Secure AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-secure

	Secure bool `json:"Secure"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksApp_EnvironmentVariable) AWSCloudFormationType() string {
	return "AWS::OpsWorks::App.EnvironmentVariable"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksApp_EnvironmentVariable) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksApp_EnvironmentVariableResources retrieves all AWSOpsWorksApp_EnvironmentVariable items from a CloudFormation template
func GetAllAWSOpsWorksApp_EnvironmentVariable(template *Template) map[string]*AWSOpsWorksApp_EnvironmentVariable {

	results := map[string]*AWSOpsWorksApp_EnvironmentVariable{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksApp_EnvironmentVariable{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksApp_EnvironmentVariableWithName retrieves all AWSOpsWorksApp_EnvironmentVariable items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksApp_EnvironmentVariable(name string, template *Template) (*AWSOpsWorksApp_EnvironmentVariable, error) {

	result := &AWSOpsWorksApp_EnvironmentVariable{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksApp_EnvironmentVariable{}, errors.New("resource not found")

}
