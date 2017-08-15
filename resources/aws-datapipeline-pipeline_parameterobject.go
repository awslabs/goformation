package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DataPipeline::Pipeline.ParameterObject AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html
type AWSDataPipelinePipeline_ParameterObject struct {

	// Attributes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html#cfn-datapipeline-pipeline-parameterobjects-attributes
	Attributes []AWSDataPipelinePipeline_ParameterAttribute `json:"Attributes"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html#cfn-datapipeline-pipeline-parameterobject-id
	Id string `json:"Id"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_ParameterObject) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.ParameterObject"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipeline_ParameterObject) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDataPipelinePipeline_ParameterObjectResources retrieves all AWSDataPipelinePipeline_ParameterObject items from a CloudFormation template
func GetAllAWSDataPipelinePipeline_ParameterObject(template *Template) map[string]*AWSDataPipelinePipeline_ParameterObject {

	results := map[string]*AWSDataPipelinePipeline_ParameterObject{}
	for name, resource := range template.Resources {
		result := &AWSDataPipelinePipeline_ParameterObject{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDataPipelinePipeline_ParameterObjectWithName retrieves all AWSDataPipelinePipeline_ParameterObject items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDataPipelinePipeline_ParameterObject(name string, template *Template) (*AWSDataPipelinePipeline_ParameterObject, error) {

	result := &AWSDataPipelinePipeline_ParameterObject{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDataPipelinePipeline_ParameterObject{}, errors.New("resource not found")

}
