package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.PortMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html
type AWSECSTaskDefinition_PortMapping struct {

	// ContainerPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html#cfn-ecs-taskdefinition-containerdefinition-portmappings-containerport

	ContainerPort int64 `json:"ContainerPort"`

	// HostPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html#cfn-ecs-taskdefinition-containerdefinition-portmappings-readonly

	HostPort int64 `json:"HostPort"`

	// Protocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-portmappings.html#cfn-ecs-taskdefinition-containerdefinition-portmappings-sourcevolume

	Protocol string `json:"Protocol"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_PortMapping) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.PortMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_PortMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_PortMappingResources retrieves all AWSECSTaskDefinition_PortMapping items from a CloudFormation template
func GetAllAWSECSTaskDefinition_PortMapping(template *Template) map[string]*AWSECSTaskDefinition_PortMapping {

	results := map[string]*AWSECSTaskDefinition_PortMapping{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_PortMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_PortMappingWithName retrieves all AWSECSTaskDefinition_PortMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_PortMapping(name string, template *Template) (*AWSECSTaskDefinition_PortMapping, error) {

	result := &AWSECSTaskDefinition_PortMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_PortMapping{}, errors.New("resource not found")

}
