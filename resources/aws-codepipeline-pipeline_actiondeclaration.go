package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.ActionDeclaration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html
type AWSCodePipelinePipeline_ActionDeclaration struct {

	// ActionTypeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid

	ActionTypeId AWSCodePipelinePipeline_ActionTypeId `json:"ActionTypeId"`

	// Configuration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-configuration

	Configuration interface{} `json:"Configuration"`

	// InputArtifacts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-inputartifacts

	InputArtifacts []AWSCodePipelinePipeline_InputArtifact `json:"InputArtifacts"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-name

	Name string `json:"Name"`

	// OutputArtifacts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-outputartifacts

	OutputArtifacts []AWSCodePipelinePipeline_OutputArtifact `json:"OutputArtifacts"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-rolearn

	RoleArn string `json:"RoleArn"`

	// RunOrder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-runorder

	RunOrder int64 `json:"RunOrder"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_ActionDeclaration) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.ActionDeclaration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_ActionDeclaration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_ActionDeclarationResources retrieves all AWSCodePipelinePipeline_ActionDeclaration items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_ActionDeclaration(template *Template) map[string]*AWSCodePipelinePipeline_ActionDeclaration {

	results := map[string]*AWSCodePipelinePipeline_ActionDeclaration{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_ActionDeclaration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_ActionDeclarationWithName retrieves all AWSCodePipelinePipeline_ActionDeclaration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_ActionDeclaration(name string, template *Template) (*AWSCodePipelinePipeline_ActionDeclaration, error) {

	result := &AWSCodePipelinePipeline_ActionDeclaration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_ActionDeclaration{}, errors.New("resource not found")

}
