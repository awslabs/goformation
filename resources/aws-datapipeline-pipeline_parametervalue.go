package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DataPipeline::Pipeline.ParameterValue AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html
type AWSDataPipelinePipeline_ParameterValue struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html#cfn-datapipeline-pipeline-parametervalues-id

	Id string `json:"Id"`

	// StringValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parametervalues.html#cfn-datapipeline-pipeline-parametervalues-stringvalue

	StringValue string `json:"StringValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_ParameterValue) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.ParameterValue"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline_ParameterValue) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipeline_ParameterValueResources retrieves all AWSDataPipelinePipeline_ParameterValue items from a CloudFormation template
func GetAllAWSDataPipelinePipeline_ParameterValue(template *Template) map[string]*AWSDataPipelinePipeline_ParameterValue {

	results := map[string]*AWSDataPipelinePipeline_ParameterValue{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline_ParameterValue{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipeline_ParameterValueWithName retrieves all AWSDataPipelinePipeline_ParameterValue items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDataPipelinePipeline_ParameterValue(name string, template *Template) (*AWSDataPipelinePipeline_ParameterValue, error) {

	result := &AWSDataPipelinePipeline_ParameterValue{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline_ParameterValue{}, errors.New("resource not found")

}
