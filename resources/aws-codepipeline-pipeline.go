package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCodePipelinePipeline AWS CloudFormation Resource (AWS::CodePipeline::Pipeline)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
type AWSCodePipelinePipeline struct {

	// ArtifactStore AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstore
	ArtifactStore AWSCodePipelinePipeline_ArtifactStore `json:"ArtifactStore"`

	// DisableInboundStageTransitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-disableinboundstagetransitions
	DisableInboundStageTransitions []AWSCodePipelinePipeline_StageTransition `json:"DisableInboundStageTransitions"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-name
	Name string `json:"Name"`

	// RestartExecutionOnUpdate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-restartexecutiononupdate
	RestartExecutionOnUpdate bool `json:"RestartExecutionOnUpdate"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-rolearn
	RoleArn string `json:"RoleArn"`

	// Stages AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-stages
	Stages []AWSCodePipelinePipeline_StageDeclaration `json:"Stages"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipelineResources retrieves all AWSCodePipelinePipeline items from a CloudFormation template
func GetAllAWSCodePipelinePipelineResources(template *Template) map[string]*AWSCodePipelinePipeline {

	results := map[string]*AWSCodePipelinePipeline{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipelineWithName retrieves all AWSCodePipelinePipeline items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCodePipelinePipelineWithName(name string, template *Template) (*AWSCodePipelinePipeline, error) {

	result := &AWSCodePipelinePipeline{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline{}, errors.New("resource not found")

}
