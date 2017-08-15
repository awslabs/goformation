package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DataPipeline::Pipeline.PipelineTag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html
type AWSDataPipelinePipeline_PipelineTag struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html#cfn-datapipeline-pipeline-pipelinetags-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html#cfn-datapipeline-pipeline-pipelinetags-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_PipelineTag) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.PipelineTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline_PipelineTag) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipeline_PipelineTagResources retrieves all AWSDataPipelinePipeline_PipelineTag items from a CloudFormation template
func GetAllAWSDataPipelinePipeline_PipelineTag(template *Template) map[string]*AWSDataPipelinePipeline_PipelineTag {

	results := map[string]*AWSDataPipelinePipeline_PipelineTag{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline_PipelineTag{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipeline_PipelineTagWithName retrieves all AWSDataPipelinePipeline_PipelineTag items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDataPipelinePipeline_PipelineTag(name string, template *Template) (*AWSDataPipelinePipeline_PipelineTag, error) {

	result := &AWSDataPipelinePipeline_PipelineTag{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline_PipelineTag{}, errors.New("resource not found")

}
