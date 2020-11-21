package databrew

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Job_Output AWS CloudFormation Resource (AWS::DataBrew::Job.Output)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html
type Job_Output struct {

	// CompressionFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-compressionformat
	CompressionFormat string `json:"CompressionFormat,omitempty"`

	// Format AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-format
	Format string `json:"Format,omitempty"`

	// Location AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-location
	Location *Job_S3Location `json:"Location,omitempty"`

	// Overwrite AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-overwrite
	Overwrite bool `json:"Overwrite,omitempty"`

	// PartitionColumns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-output.html#cfn-databrew-job-output-partitioncolumns
	PartitionColumns []string `json:"PartitionColumns,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Job_Output) AWSCloudFormationType() string {
	return "AWS::DataBrew::Job.Output"
}
