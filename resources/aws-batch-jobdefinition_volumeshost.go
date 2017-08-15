package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.VolumesHost AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html
type AWSBatchJobDefinition_VolumesHost struct {

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html#cfn-batch-jobdefinition-volumeshost-sourcepath
	SourcePath string `json:"SourcePath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_VolumesHost) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.VolumesHost"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_VolumesHost) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
