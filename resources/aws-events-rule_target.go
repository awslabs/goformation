package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Events::Rule.Target AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html
type AWSEventsRule_Target struct {

	// Arn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-arn
	Arn string `json:"Arn"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-id
	Id string `json:"Id"`

	// Input AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-input
	Input string `json:"Input"`

	// InputPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-inputpath
	InputPath string `json:"InputPath"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-rolearn
	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_Target) AWSCloudFormationType() string {
	return "AWS::Events::Rule.Target"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEventsRule_Target) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEventsRule_TargetResources retrieves all AWSEventsRule_Target items from a CloudFormation template
func GetAllAWSEventsRule_Target(template *Template) map[string]*AWSEventsRule_Target {

	results := map[string]*AWSEventsRule_Target{}
	for name, resource := range template.Resources {
		result := &AWSEventsRule_Target{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEventsRule_TargetWithName retrieves all AWSEventsRule_Target items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEventsRule_Target(name string, template *Template) (*AWSEventsRule_Target, error) {

	result := &AWSEventsRule_Target{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEventsRule_Target{}, errors.New("resource not found")

}
