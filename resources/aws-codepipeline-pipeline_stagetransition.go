package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.StageTransition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-disableinboundstagetransitions.html
type AWSCodePipelinePipeline_StageTransition struct {

	// Reason AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-disableinboundstagetransitions.html#cfn-codepipeline-pipeline-disableinboundstagetransitions-reason

	Reason string `json:"Reason"`

	// StageName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-disableinboundstagetransitions.html#cfn-codepipeline-pipeline-disableinboundstagetransitions-stagename

	StageName string `json:"StageName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_StageTransition) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.StageTransition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_StageTransition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_StageTransitionResources retrieves all AWSCodePipelinePipeline_StageTransition items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_StageTransition(template *Template) map[string]*AWSCodePipelinePipeline_StageTransition {

	results := map[string]*AWSCodePipelinePipeline_StageTransition{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_StageTransition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_StageTransitionWithName retrieves all AWSCodePipelinePipeline_StageTransition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_StageTransition(name string, template *Template) (*AWSCodePipelinePipeline_StageTransition, error) {

	result := &AWSCodePipelinePipeline_StageTransition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_StageTransition{}, errors.New("resource not found")

}
