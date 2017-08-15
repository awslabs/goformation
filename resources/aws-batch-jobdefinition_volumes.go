package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.Volumes AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html
type AWSBatchJobDefinition_Volumes struct {

	// Host AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-host
	Host AWSBatchJobDefinition_VolumesHost `json:"Host"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_Volumes) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.Volumes"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_Volumes) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
