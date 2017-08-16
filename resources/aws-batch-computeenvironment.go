package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSBatchComputeEnvironment AWS CloudFormation Resource (AWS::Batch::ComputeEnvironment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html
type AWSBatchComputeEnvironment struct {

	// ComputeEnvironmentName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-computeenvironmentname
	ComputeEnvironmentName string `json:"ComputeEnvironmentName"`

	// ComputeResources AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-computeresources
	ComputeResources AWSBatchComputeEnvironment_ComputeResources `json:"ComputeResources"`

	// ServiceRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-servicerole
	ServiceRole string `json:"ServiceRole"`

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-state
	State string `json:"State"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html#cfn-batch-computeenvironment-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchComputeEnvironment) AWSCloudFormationType() string {
	return "AWS::Batch::ComputeEnvironment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchComputeEnvironment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchComputeEnvironmentResources retrieves all AWSBatchComputeEnvironment items from a CloudFormation template
func (t *Template) GetAllAWSBatchComputeEnvironmentResources() map[string]*AWSBatchComputeEnvironment {

	results := map[string]*AWSBatchComputeEnvironment{}
	for name, resource := range t.Resources {
		result := &AWSBatchComputeEnvironment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchComputeEnvironmentWithName retrieves all AWSBatchComputeEnvironment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBatchComputeEnvironmentWithName(name string) (*AWSBatchComputeEnvironment, error) {

	result := &AWSBatchComputeEnvironment{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchComputeEnvironment{}, errors.New("resource not found")

}
