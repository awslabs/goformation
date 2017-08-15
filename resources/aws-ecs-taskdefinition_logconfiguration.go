package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.LogConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html
type AWSECSTaskDefinition_LogConfiguration struct {

	// LogDriver AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration-logdriver
	LogDriver string `json:"LogDriver"`

	// Options AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-logconfiguration.html#cfn-ecs-taskdefinition-containerdefinition-logconfiguration-options
	Options map[string]string `json:"Options"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_LogConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.LogConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_LogConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_LogConfigurationResources retrieves all AWSECSTaskDefinition_LogConfiguration items from a CloudFormation template
func GetAllAWSECSTaskDefinition_LogConfiguration(template *Template) map[string]*AWSECSTaskDefinition_LogConfiguration {

	results := map[string]*AWSECSTaskDefinition_LogConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_LogConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_LogConfigurationWithName retrieves all AWSECSTaskDefinition_LogConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_LogConfiguration(name string, template *Template) (*AWSECSTaskDefinition_LogConfiguration, error) {

	result := &AWSECSTaskDefinition_LogConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_LogConfiguration{}, errors.New("resource not found")

}
