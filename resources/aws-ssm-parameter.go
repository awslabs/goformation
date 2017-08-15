package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SSM::Parameter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html
type AWSSSMParameter struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-description
	Description string `json:"Description"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-name
	Name string `json:"Name"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html#cfn-ssm-parameter-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMParameter) AWSCloudFormationType() string {
	return "AWS::SSM::Parameter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMParameter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSSMParameterResources retrieves all AWSSSMParameter items from a CloudFormation template
func GetAllAWSSSMParameterResources(template *Template) map[string]*AWSSSMParameter {

	results := map[string]*AWSSSMParameter{}
	for name, resource := range template.Resources {
		result := &AWSSSMParameter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSSMParameterWithName retrieves all AWSSSMParameter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSSSMParameterWithName(name string, template *Template) (*AWSSSMParameter, error) {

	result := &AWSSSMParameter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSSMParameter{}, errors.New("resource not found")

}
