package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.MountPoint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html
type AWSECSTaskDefinition_MountPoint struct {

	// ContainerPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-containerpath
	ContainerPath string `json:"ContainerPath"`

	// ReadOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-readonly
	ReadOnly bool `json:"ReadOnly"`

	// SourceVolume AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-sourcevolume
	SourceVolume string `json:"SourceVolume"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_MountPoint) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.MountPoint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_MountPoint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_MountPointResources retrieves all AWSECSTaskDefinition_MountPoint items from a CloudFormation template
func GetAllAWSECSTaskDefinition_MountPoint(template *Template) map[string]*AWSECSTaskDefinition_MountPoint {

	results := map[string]*AWSECSTaskDefinition_MountPoint{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_MountPoint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_MountPointWithName retrieves all AWSECSTaskDefinition_MountPoint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_MountPoint(name string, template *Template) (*AWSECSTaskDefinition_MountPoint, error) {

	result := &AWSECSTaskDefinition_MountPoint{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_MountPoint{}, errors.New("resource not found")

}
