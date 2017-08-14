package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DataPipeline::Pipeline.ParameterAttribute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html
type AWSDataPipelinePipeline_ParameterAttribute struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-key

	Key string `json:"Key"`

	// StringValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-stringvalue

	StringValue string `json:"StringValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_ParameterAttribute) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.ParameterAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline_ParameterAttribute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipeline_ParameterAttributeResources retrieves all AWSDataPipelinePipeline_ParameterAttribute items from a CloudFormation template
func GetAllAWSDataPipelinePipeline_ParameterAttribute(template *Template) map[string]*AWSDataPipelinePipeline_ParameterAttribute {

	results := map[string]*AWSDataPipelinePipeline_ParameterAttribute{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline_ParameterAttribute{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipeline_ParameterAttributeWithName retrieves all AWSDataPipelinePipeline_ParameterAttribute items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDataPipelinePipeline_ParameterAttribute(name string, template *Template) (*AWSDataPipelinePipeline_ParameterAttribute, error) {

	result := &AWSDataPipelinePipeline_ParameterAttribute{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline_ParameterAttribute{}, errors.New("resource not found")

}
