package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.KeyValuePair AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html
type AWSECSTaskDefinition_KeyValuePair struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html#cfn-ecs-taskdefinition-containerdefinition-environment-name

	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-environment.html#cfn-ecs-taskdefinition-containerdefinition-environment-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_KeyValuePair) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.KeyValuePair"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_KeyValuePair) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_KeyValuePairResources retrieves all AWSECSTaskDefinition_KeyValuePair items from a CloudFormation template
func GetAllAWSECSTaskDefinition_KeyValuePair(template *Template) map[string]*AWSECSTaskDefinition_KeyValuePair {

	results := map[string]*AWSECSTaskDefinition_KeyValuePair{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_KeyValuePair{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_KeyValuePairWithName retrieves all AWSECSTaskDefinition_KeyValuePair items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_KeyValuePair(name string, template *Template) (*AWSECSTaskDefinition_KeyValuePair, error) {

	result := &AWSECSTaskDefinition_KeyValuePair{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_KeyValuePair{}, errors.New("resource not found")

}
