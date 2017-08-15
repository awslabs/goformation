package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobQueue.ComputeEnvironmentOrder AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html
type AWSBatchJobQueue_ComputeEnvironmentOrder struct {

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
func (r *AWSBatchJobQueue_ComputeEnvironmentOrder) AWSCloudFormationType() string {
	return "AWS::Batch::JobQueue.ComputeEnvironmentOrder"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobQueue_ComputeEnvironmentOrder) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobQueue_ComputeEnvironmentOrderResources retrieves all AWSBatchJobQueue_ComputeEnvironmentOrder items from a CloudFormation template
func GetAllAWSBatchJobQueue_ComputeEnvironmentOrder(template *Template) map[string]*AWSBatchJobQueue_ComputeEnvironmentOrder {

	results := map[string]*AWSBatchJobQueue_ComputeEnvironmentOrder{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobQueue_ComputeEnvironmentOrder{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobQueue_ComputeEnvironmentOrderWithName retrieves all AWSBatchJobQueue_ComputeEnvironmentOrder items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobQueue_ComputeEnvironmentOrder(name string, template *Template) (*AWSBatchJobQueue_ComputeEnvironmentOrder, error) {

	result := &AWSBatchJobQueue_ComputeEnvironmentOrder{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobQueue_ComputeEnvironmentOrder{}, errors.New("resource not found")

}
