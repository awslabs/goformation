// Code generated by "go generate". Please don't change this file directly.

package codepipeline

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Pipeline_GitFilePathFilterCriteria AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.GitFilePathFilterCriteria)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-file-path-filter-criteria.html
type Pipeline_GitFilePathFilterCriteria struct {

	// Excludes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-file-path-filter-criteria.html#aws-properties-codepipeline-pipeline-triggers-git-file-path-pattern
	Excludes []string `json:"Excludes,omitempty"`

	// Includes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-triggers-git-file-path-filter-criteria.html#aws-properties-codepipeline-pipeline-triggers-git-file-path-pattern
	Includes []string `json:"Includes,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Pipeline_GitFilePathFilterCriteria) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.GitFilePathFilterCriteria"
}