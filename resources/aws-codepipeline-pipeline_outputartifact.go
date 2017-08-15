package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.OutputArtifact AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-outputartifacts.html
type AWSCodePipelinePipeline_OutputArtifact struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-outputartifacts.html#cfn-codepipeline-pipeline-stages-actions-outputartifacts-name

	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_OutputArtifact) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.OutputArtifact"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_OutputArtifact) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_OutputArtifactResources retrieves all AWSCodePipelinePipeline_OutputArtifact items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_OutputArtifact(template *Template) map[string]*AWSCodePipelinePipeline_OutputArtifact {

	results := map[string]*AWSCodePipelinePipeline_OutputArtifact{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_OutputArtifact{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_OutputArtifactWithName retrieves all AWSCodePipelinePipeline_OutputArtifact items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_OutputArtifact(name string, template *Template) (*AWSCodePipelinePipeline_OutputArtifact, error) {

	result := &AWSCodePipelinePipeline_OutputArtifact{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_OutputArtifact{}, errors.New("resource not found")

}
