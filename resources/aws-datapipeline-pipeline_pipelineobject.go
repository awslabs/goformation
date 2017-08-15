package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DataPipeline::Pipeline.PipelineObject AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html
type AWSDataPipelinePipeline_PipelineObject struct {

	// Fields AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-fields
	Fields []AWSDataPipelinePipeline_Field `json:"Fields"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-id
	Id string `json:"Id"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_PipelineObject) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.PipelineObject"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline_PipelineObject) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipeline_PipelineObjectResources retrieves all AWSDataPipelinePipeline_PipelineObject items from a CloudFormation template
func GetAllAWSDataPipelinePipeline_PipelineObject(template *Template) map[string]*AWSDataPipelinePipeline_PipelineObject {

	results := map[string]*AWSDataPipelinePipeline_PipelineObject{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline_PipelineObject{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipeline_PipelineObjectWithName retrieves all AWSDataPipelinePipeline_PipelineObject items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDataPipelinePipeline_PipelineObject(name string, template *Template) (*AWSDataPipelinePipeline_PipelineObject, error) {

	result := &AWSDataPipelinePipeline_PipelineObject{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline_PipelineObject{}, errors.New("resource not found")

}
