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
