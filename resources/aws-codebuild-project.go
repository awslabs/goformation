package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeBuild::Project AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
type AWSCodeBuildProject struct {

	// Artifacts AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-artifacts
	Artifacts AWSCodeBuildProject_Artifacts `json:"Artifacts"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-description
	Description string `json:"Description"`

	// EncryptionKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-encryptionkey
	EncryptionKey string `json:"EncryptionKey"`

	// Environment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment
	Environment AWSCodeBuildProject_Environment `json:"Environment"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-name
	Name string `json:"Name"`

	// ServiceRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-servicerole
	ServiceRole string `json:"ServiceRole"`

	// Source AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-source
	Source AWSCodeBuildProject_Source `json:"Source"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-tags
	Tags []Tag `json:"Tags"`

	// TimeoutInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-timeoutinminutes
	TimeoutInMinutes int64 `json:"TimeoutInMinutes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProject) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeBuildProjectResources retrieves all AWSCodeBuildProject items from a CloudFormation template
func GetAllAWSCodeBuildProjectResources(template *Template) map[string]*AWSCodeBuildProject {

	results := map[string]*AWSCodeBuildProject{}
	for name, resource := range template.Resources {
		result := &AWSCodeBuildProject{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeBuildProjectWithName retrieves all AWSCodeBuildProject items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCodeBuildProjectWithName(name string, template *Template) (*AWSCodeBuildProject, error) {

	result := &AWSCodeBuildProject{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeBuildProject{}, errors.New("resource not found")

}
