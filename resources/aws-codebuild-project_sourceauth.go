package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeBuild::Project.SourceAuth AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html
type AWSCodeBuildProject_SourceAuth struct {

	// Resource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-resource
	Resource string `json:"Resource"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_SourceAuth) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.SourceAuth"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProject_SourceAuth) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeBuildProject_SourceAuthResources retrieves all AWSCodeBuildProject_SourceAuth items from a CloudFormation template
func GetAllAWSCodeBuildProject_SourceAuth(template *Template) map[string]*AWSCodeBuildProject_SourceAuth {

	results := map[string]*AWSCodeBuildProject_SourceAuth{}
	for name, resource := range template.Resources {
		result := &AWSCodeBuildProject_SourceAuth{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeBuildProject_SourceAuthWithName retrieves all AWSCodeBuildProject_SourceAuth items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeBuildProject_SourceAuth(name string, template *Template) (*AWSCodeBuildProject_SourceAuth, error) {

	result := &AWSCodeBuildProject_SourceAuth{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeBuildProject_SourceAuth{}, errors.New("resource not found")

}
