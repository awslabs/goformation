package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.RetryStrategy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html
type AWSBatchJobDefinition_RetryStrategy struct {

	// Attempts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html#cfn-batch-jobdefinition-retrystrategy-attempts
	Attempts int64 `json:"Attempts"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_RetryStrategy) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.RetryStrategy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_RetryStrategy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
