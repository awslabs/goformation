package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeBuild::Project.EnvironmentVariable AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html
type AWSCodeBuildProject_EnvironmentVariable struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-name

	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_EnvironmentVariable) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.EnvironmentVariable"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProject_EnvironmentVariable) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeBuildProject_EnvironmentVariableResources retrieves all AWSCodeBuildProject_EnvironmentVariable items from a CloudFormation template
func GetAllAWSCodeBuildProject_EnvironmentVariable(template *Template) map[string]*AWSCodeBuildProject_EnvironmentVariable {

	results := map[string]*AWSCodeBuildProject_EnvironmentVariable{}
	for name, resource := range template.Resources {
		result := &AWSCodeBuildProject_EnvironmentVariable{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeBuildProject_EnvironmentVariableWithName retrieves all AWSCodeBuildProject_EnvironmentVariable items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeBuildProject_EnvironmentVariable(name string, template *Template) (*AWSCodeBuildProject_EnvironmentVariable, error) {

	result := &AWSCodeBuildProject_EnvironmentVariable{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeBuildProject_EnvironmentVariable{}, errors.New("resource not found")

}
