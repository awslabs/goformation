package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIoTPolicy AWS CloudFormation Resource (AWS::IoT::Policy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html
type AWSIoTPolicy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument"`

	// PolicyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policyname
	PolicyName string `json:"PolicyName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTPolicy) AWSCloudFormationType() string {
	return "AWS::IoT::Policy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTPolicyResources retrieves all AWSIoTPolicy items from a CloudFormation template
func (t *Template) GetAllAWSIoTPolicyResources() map[string]*AWSIoTPolicy {

	results := map[string]*AWSIoTPolicy{}
	for name, resource := range t.Resources {
		result := &AWSIoTPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTPolicyWithName retrieves all AWSIoTPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTPolicyWithName(name string) (*AWSIoTPolicy, error) {

	result := &AWSIoTPolicy{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTPolicy{}, errors.New("resource not found")

}
