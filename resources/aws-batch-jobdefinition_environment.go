package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.Environment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html
type AWSBatchJobDefinition_Environment struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html#cfn-batch-jobdefinition-environment-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html#cfn-batch-jobdefinition-environment-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_Environment) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_Environment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
