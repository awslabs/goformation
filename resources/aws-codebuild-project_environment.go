package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeBuild::Project.Environment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html
type AWSCodeBuildProject_Environment struct {

	// ComputeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-computetype
	ComputeType string `json:"ComputeType"`

	// EnvironmentVariables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-environmentvariables
	EnvironmentVariables []AWSCodeBuildProject_EnvironmentVariable `json:"EnvironmentVariables"`

	// Image AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-image
	Image string `json:"Image"`

	// PrivilegedMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-privilegedmode
	PrivilegedMode bool `json:"PrivilegedMode"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html#cfn-codebuild-project-environment-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_Environment) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProject_Environment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeBuildProject_EnvironmentResources retrieves all AWSCodeBuildProject_Environment items from a CloudFormation template
func GetAllAWSCodeBuildProject_Environment(template *Template) map[string]*AWSCodeBuildProject_Environment {

	results := map[string]*AWSCodeBuildProject_Environment{}
	for name, resource := range template.Resources {
		result := &AWSCodeBuildProject_Environment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeBuildProject_EnvironmentWithName retrieves all AWSCodeBuildProject_Environment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeBuildProject_Environment(name string, template *Template) (*AWSCodeBuildProject_Environment, error) {

	result := &AWSCodeBuildProject_Environment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeBuildProject_Environment{}, errors.New("resource not found")

}
