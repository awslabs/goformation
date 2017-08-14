package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SQS::Queue AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html
type AWSSQSQueue struct {

	// ContentBasedDeduplication AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-contentbaseddeduplication

	ContentBasedDeduplication bool `json:"ContentBasedDeduplication"`

	// DelaySeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-delayseconds

	DelaySeconds int64 `json:"DelaySeconds"`

	// FifoQueue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-fifoqueue

	FifoQueue bool `json:"FifoQueue"`

	// MaximumMessageSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-maxmesgsize

	MaximumMessageSize int64 `json:"MaximumMessageSize"`

	// MessageRetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-msgretentionperiod

	MessageRetentionPeriod int64 `json:"MessageRetentionPeriod"`

	// QueueName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-name

	QueueName string `json:"QueueName"`

	// ReceiveMessageWaitTimeSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-receivemsgwaittime

	ReceiveMessageWaitTimeSeconds int64 `json:"ReceiveMessageWaitTimeSeconds"`

	// RedrivePolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-redrive

	RedrivePolicy interface{} `json:"RedrivePolicy"`

	// VisibilityTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-visiblitytimeout

	VisibilityTimeout int64 `json:"VisibilityTimeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSQSQueue) AWSCloudFormationType() string {
	return "AWS::SQS::Queue"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSQSQueue) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSQSQueueResources retrieves all AWSSQSQueue items from a CloudFormation template
func GetAllAWSSQSQueue(template *Template) map[string]*AWSSQSQueue {

	results := map[string]*AWSSQSQueue{}
	for name, resource := range template.Resources {
		result := &AWSSQSQueue{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSQSQueueWithName retrieves all AWSSQSQueue items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSQSQueue(name string, template *Template) (*AWSSQSQueue, error) {

	result := &AWSSQSQueue{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSQSQueue{}, errors.New("resource not found")

}
