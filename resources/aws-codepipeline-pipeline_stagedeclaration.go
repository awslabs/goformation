package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.StageDeclaration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html
type AWSCodePipelinePipeline_StageDeclaration struct {

	// Actions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html#cfn-codepipeline-pipeline-stages-actions
	Actions []AWSCodePipelinePipeline_ActionDeclaration `json:"Actions"`

	// Blockers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html#cfn-codepipeline-pipeline-stages-blockers
	Blockers []AWSCodePipelinePipeline_BlockerDeclaration `json:"Blockers"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages.html#cfn-codepipeline-pipeline-stages-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_StageDeclaration) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.StageDeclaration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_StageDeclaration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_StageDeclarationResources retrieves all AWSCodePipelinePipeline_StageDeclaration items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_StageDeclaration(template *Template) map[string]*AWSCodePipelinePipeline_StageDeclaration {

	results := map[string]*AWSCodePipelinePipeline_StageDeclaration{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_StageDeclaration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_StageDeclarationWithName retrieves all AWSCodePipelinePipeline_StageDeclaration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_StageDeclaration(name string, template *Template) (*AWSCodePipelinePipeline_StageDeclaration, error) {

	result := &AWSCodePipelinePipeline_StageDeclaration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_StageDeclaration{}, errors.New("resource not found")

}
