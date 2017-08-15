package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SQS::QueuePolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
type AWSSQSQueuePolicy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-policydoc
	PolicyDocument interface{} `json:"PolicyDocument"`

	// Queues AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-queues
	Queues []string `json:"Queues"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSQSQueuePolicy) AWSCloudFormationType() string {
	return "AWS::SQS::QueuePolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSQSQueuePolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSQSQueuePolicyResources retrieves all AWSSQSQueuePolicy items from a CloudFormation template
func GetAllAWSSQSQueuePolicy(template *Template) map[string]*AWSSQSQueuePolicy {

	results := map[string]*AWSSQSQueuePolicy{}
	for name, resource := range template.Resources {
		result := &AWSSQSQueuePolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSQSQueuePolicyWithName retrieves all AWSSQSQueuePolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSQSQueuePolicy(name string, template *Template) (*AWSSQSQueuePolicy, error) {

	result := &AWSSQSQueuePolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSQSQueuePolicy{}, errors.New("resource not found")

}
