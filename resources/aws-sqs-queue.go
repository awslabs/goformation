package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSSQSQueue AWS CloudFormation Resource (AWS::SQS::Queue)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html
type AWSSQSQueue struct {

	// ContentBasedDeduplication AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-contentbaseddeduplication
	ContentBasedDeduplication bool `json:"ContentBasedDeduplication"`

	// DelaySeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-delayseconds
	DelaySeconds int `json:"DelaySeconds"`

	// FifoQueue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-fifoqueue
	FifoQueue bool `json:"FifoQueue"`

	// MaximumMessageSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-maxmesgsize
	MaximumMessageSize int `json:"MaximumMessageSize"`

	// MessageRetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-msgretentionperiod
	MessageRetentionPeriod int `json:"MessageRetentionPeriod"`

	// QueueName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-name
	QueueName string `json:"QueueName"`

	// ReceiveMessageWaitTimeSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-receivemsgwaittime
	ReceiveMessageWaitTimeSeconds int `json:"ReceiveMessageWaitTimeSeconds"`

	// RedrivePolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-redrive
	RedrivePolicy interface{} `json:"RedrivePolicy"`

	// VisibilityTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html#aws-sqs-queue-visiblitytimeout
	VisibilityTimeout int `json:"VisibilityTimeout"`
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
func (t *Template) GetAllAWSSQSQueueResources() map[string]*AWSSQSQueue {

	results := map[string]*AWSSQSQueue{}
	for name, resource := range t.Resources {
		result := &AWSSQSQueue{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSQSQueueWithName retrieves all AWSSQSQueue items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSQSQueueWithName(name string) (*AWSSQSQueue, error) {

	result := &AWSSQSQueue{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSQSQueue{}, errors.New("resource not found")

}
