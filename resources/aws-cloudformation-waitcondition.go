package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCloudFormationWaitCondition AWS CloudFormation Resource (AWS::CloudFormation::WaitCondition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html
type AWSCloudFormationWaitCondition struct {

	// Count AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-count
	Count int `json:"Count"`

	// Handle AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-handle
	Handle string `json:"Handle"`

	// Timeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html#cfn-waitcondition-timeout
	Timeout string `json:"Timeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationWaitCondition) AWSCloudFormationType() string {
	return "AWS::CloudFormation::WaitCondition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationWaitCondition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFormationWaitConditionResources retrieves all AWSCloudFormationWaitCondition items from a CloudFormation template
func (t *Template) GetAllAWSCloudFormationWaitConditionResources() map[string]*AWSCloudFormationWaitCondition {

	results := map[string]*AWSCloudFormationWaitCondition{}
	for name, resource := range t.Resources {
		result := &AWSCloudFormationWaitCondition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFormationWaitConditionWithName retrieves all AWSCloudFormationWaitCondition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationWaitConditionWithName(name string) (*AWSCloudFormationWaitCondition, error) {

	result := &AWSCloudFormationWaitCondition{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFormationWaitCondition{}, errors.New("resource not found")

}
