package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSDataPipelinePipeline AWS CloudFormation Resource (AWS::DataPipeline::Pipeline)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html
type AWSDataPipelinePipeline struct {

	// Activate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-activate
	Activate bool `json:"Activate"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-description
	Description string `json:"Description"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-name
	Name string `json:"Name"`

	// ParameterObjects AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parameterobjects
	ParameterObjects []AWSDataPipelinePipeline_ParameterObject `json:"ParameterObjects"`

	// ParameterValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parametervalues
	ParameterValues []AWSDataPipelinePipeline_ParameterValue `json:"ParameterValues"`

	// PipelineObjects AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelineobjects
	PipelineObjects []AWSDataPipelinePipeline_PipelineObject `json:"PipelineObjects"`

	// PipelineTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelinetags
	PipelineTags []AWSDataPipelinePipeline_PipelineTag `json:"PipelineTags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipelineResources retrieves all AWSDataPipelinePipeline items from a CloudFormation template
func GetAllAWSDataPipelinePipelineResources(template *Template) map[string]*AWSDataPipelinePipeline {

	results := map[string]*AWSDataPipelinePipeline{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipelineWithName retrieves all AWSDataPipelinePipeline items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSDataPipelinePipelineWithName(name string, template *Template) (*AWSDataPipelinePipeline, error) {

	result := &AWSDataPipelinePipeline{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline{}, errors.New("resource not found")

}
