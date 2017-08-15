package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.InputArtifact AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-inputartifacts.html
type AWSCodePipelinePipeline_InputArtifact struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-inputartifacts.html#cfn-codepipeline-pipeline-stages-actions-inputartifacts-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_InputArtifact) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.InputArtifact"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_InputArtifact) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_InputArtifactResources retrieves all AWSCodePipelinePipeline_InputArtifact items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_InputArtifact(template *Template) map[string]*AWSCodePipelinePipeline_InputArtifact {

	results := map[string]*AWSCodePipelinePipeline_InputArtifact{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_InputArtifact{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_InputArtifactWithName retrieves all AWSCodePipelinePipeline_InputArtifact items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_InputArtifact(name string, template *Template) (*AWSCodePipelinePipeline_InputArtifact, error) {

	result := &AWSCodePipelinePipeline_InputArtifact{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_InputArtifact{}, errors.New("resource not found")

}
