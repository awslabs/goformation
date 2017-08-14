package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DataPipeline::Pipeline.Field AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html
type AWSDataPipelinePipeline_Field struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-key

	Key string `json:"Key"`

	// RefValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-refvalue

	RefValue string `json:"RefValue"`

	// StringValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-stringvalue

	StringValue string `json:"StringValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_Field) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.Field"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline_Field) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipeline_FieldResources retrieves all AWSDataPipelinePipeline_Field items from a CloudFormation template
func GetAllAWSDataPipelinePipeline_Field(template *Template) map[string]*AWSDataPipelinePipeline_Field {

	results := map[string]*AWSDataPipelinePipeline_Field{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline_Field{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipeline_FieldWithName retrieves all AWSDataPipelinePipeline_Field items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDataPipelinePipeline_Field(name string, template *Template) (*AWSDataPipelinePipeline_Field, error) {

	result := &AWSDataPipelinePipeline_Field{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline_Field{}, errors.New("resource not found")

}
