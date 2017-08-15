package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

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

// GetAllAWSBatchJobQueueResources retrieves all AWSBatchJobQueue items from a CloudFormation template
func GetAllAWSBatchJobQueue(template *Template) map[string]*AWSBatchJobQueue {

	results := map[string]*AWSBatchJobQueue{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobQueue{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobQueueWithName retrieves all AWSBatchJobQueue items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobQueue(name string, template *Template) (*AWSBatchJobQueue, error) {

	result := &AWSBatchJobQueue{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobQueue{}, errors.New("resource not found")

}
