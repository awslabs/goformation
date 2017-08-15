package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeBuild::Project.Source AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html
type AWSCodeBuildProject_Source struct {

	// Auth AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-auth

	Auth AWSCodeBuildProject_SourceAuth `json:"Auth"`

	// BuildSpec AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-buildspec

	BuildSpec string `json:"BuildSpec"`

	// Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-location

	Location string `json:"Location"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_Source) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.Source"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProject_Source) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeBuildProject_SourceResources retrieves all AWSCodeBuildProject_Source items from a CloudFormation template
func GetAllAWSCodeBuildProject_Source(template *Template) map[string]*AWSCodeBuildProject_Source {

	results := map[string]*AWSCodeBuildProject_Source{}
	for name, resource := range template.Resources {
		result := &AWSCodeBuildProject_Source{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeBuildProject_SourceWithName retrieves all AWSCodeBuildProject_Source items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeBuildProject_Source(name string, template *Template) (*AWSCodeBuildProject_Source, error) {

	result := &AWSCodeBuildProject_Source{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeBuildProject_Source{}, errors.New("resource not found")

}
