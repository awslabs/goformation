package resources

// AWS::SQS::QueuePolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
type AWSSQSQueuePolicy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-policydoc
	PolicyDocument object `json:"PolicyDocument"`

	// Queues AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-queues
	Queues []AWSSQSQueuePolicystring `json:"Queues"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSQSQueuePolicy) AWSCloudFormationType() string {
	return "AWS::SQS::QueuePolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSQSQueuePolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
