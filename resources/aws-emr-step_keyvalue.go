package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Step.KeyValue AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html
type AWSEMRStep_KeyValue struct {

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html#cfn-elasticmapreduce-step-keyvalue-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html#cfn-elasticmapreduce-step-keyvalue-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRStep_KeyValue) AWSCloudFormationType() string {
	return "AWS::EMR::Step.KeyValue"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRStep_KeyValue) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRStep_KeyValueResources retrieves all AWSEMRStep_KeyValue items from a CloudFormation template
func GetAllAWSEMRStep_KeyValue(template *Template) map[string]*AWSEMRStep_KeyValue {

	results := map[string]*AWSEMRStep_KeyValue{}
	for name, resource := range template.Resources {
		result := &AWSEMRStep_KeyValue{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRStep_KeyValueWithName retrieves all AWSEMRStep_KeyValue items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRStep_KeyValue(name string, template *Template) (*AWSEMRStep_KeyValue, error) {

	result := &AWSEMRStep_KeyValue{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRStep_KeyValue{}, errors.New("resource not found")

}
