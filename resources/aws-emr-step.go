package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEMRStep AWS CloudFormation Resource (AWS::EMR::Step)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html
type AWSEMRStep struct {

	// ActionOnFailure AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-actiononfailure
	ActionOnFailure string `json:"ActionOnFailure"`

	// HadoopJarStep AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-hadoopjarstep
	HadoopJarStep AWSEMRStep_HadoopJarStepConfig `json:"HadoopJarStep"`

	// JobFlowId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-jobflowid
	JobFlowId string `json:"JobFlowId"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRStep) AWSCloudFormationType() string {
	return "AWS::EMR::Step"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRStep) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRStepResources retrieves all AWSEMRStep items from a CloudFormation template
func (t *Template) GetAllAWSEMRStepResources() map[string]*AWSEMRStep {

	results := map[string]*AWSEMRStep{}
	for name, resource := range t.Resources {
		result := &AWSEMRStep{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRStepWithName retrieves all AWSEMRStep items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRStepWithName(name string) (*AWSEMRStep, error) {

	result := &AWSEMRStep{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRStep{}, errors.New("resource not found")

}
