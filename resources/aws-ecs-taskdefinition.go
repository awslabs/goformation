package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSECSTaskDefinition AWS CloudFormation Resource (AWS::ECS::TaskDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
type AWSECSTaskDefinition struct {

	// ContainerDefinitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
	ContainerDefinitions []AWSECSTaskDefinition_ContainerDefinition `json:"ContainerDefinitions"`

	// Family AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
	Family string `json:"Family"`

	// NetworkMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
	NetworkMode string `json:"NetworkMode"`

	// PlacementConstraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
	PlacementConstraints []AWSECSTaskDefinition_TaskDefinitionPlacementConstraint `json:"PlacementConstraints"`

	// TaskRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
	TaskRoleArn string `json:"TaskRoleArn"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
	Volumes []AWSECSTaskDefinition_Volume `json:"Volumes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinitionResources retrieves all AWSECSTaskDefinition items from a CloudFormation template
func GetAllAWSECSTaskDefinitionResources(template *Template) map[string]*AWSECSTaskDefinition {

	results := map[string]*AWSECSTaskDefinition{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinitionWithName retrieves all AWSECSTaskDefinition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSECSTaskDefinitionWithName(name string, template *Template) (*AWSECSTaskDefinition, error) {

	result := &AWSECSTaskDefinition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition{}, errors.New("resource not found")

}
