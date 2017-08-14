package resources

// AWS::Batch::JobQueue.ComputeEnvironmentOrder AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html
type AWSBatchJobQueueComputeEnvironmentOrder struct {

	// ComputeEnvironment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html#cfn-batch-jobqueue-computeenvironmentorder-computeenvironment
	ComputeEnvironment string `json:"ComputeEnvironment"`

	// Order AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html#cfn-batch-jobqueue-computeenvironmentorder-order
	Order int64 `json:"Order"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobQueueComputeEnvironmentOrder) AWSCloudFormationType() string {
	return "AWS::Batch::JobQueue.ComputeEnvironmentOrder"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobQueueComputeEnvironmentOrder) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
