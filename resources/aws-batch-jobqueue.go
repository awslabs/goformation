package resources

// AWS::Batch::JobQueue AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html
type AWSBatchJobQueue struct {

	// ComputeEnvironmentOrder AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-computeenvironmentorder

	ComputeEnvironmentOrder []AWSBatchJobQueue_ComputeEnvironmentOrder `json:"ComputeEnvironmentOrder"`

	// JobQueueName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-jobqueuename

	JobQueueName string `json:"JobQueueName"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-priority

	Priority int64 `json:"Priority"`

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-state

	State string `json:"State"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobQueue) AWSCloudFormationType() string {
	return "AWS::Batch::JobQueue"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobQueue) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
