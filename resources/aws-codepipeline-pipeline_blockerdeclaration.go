package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.BlockerDeclaration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html
type AWSCodePipelinePipeline_BlockerDeclaration struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html#cfn-codepipeline-pipeline-stages-blockers-name

	Name string `json:"Name"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html#cfn-codepipeline-pipeline-stages-blockers-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_BlockerDeclaration) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.BlockerDeclaration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_BlockerDeclaration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_BlockerDeclarationResources retrieves all AWSCodePipelinePipeline_BlockerDeclaration items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_BlockerDeclaration(template *Template) map[string]*AWSCodePipelinePipeline_BlockerDeclaration {

	results := map[string]*AWSCodePipelinePipeline_BlockerDeclaration{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_BlockerDeclaration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_BlockerDeclarationWithName retrieves all AWSCodePipelinePipeline_BlockerDeclaration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_BlockerDeclaration(name string, template *Template) (*AWSCodePipelinePipeline_BlockerDeclaration, error) {

	result := &AWSCodePipelinePipeline_BlockerDeclaration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_BlockerDeclaration{}, errors.New("resource not found")

}
