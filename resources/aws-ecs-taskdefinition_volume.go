package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.Volume AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html
type AWSECSTaskDefinition_Volume struct {

	// Host AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html#cfn-ecs-taskdefinition-volumes-host
	Host AWSECSTaskDefinition_HostVolumeProperties `json:"Host"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes.html#cfn-ecs-taskdefinition-volumes-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_Volume) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.Volume"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_Volume) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_VolumeResources retrieves all AWSECSTaskDefinition_Volume items from a CloudFormation template
func GetAllAWSECSTaskDefinition_Volume(template *Template) map[string]*AWSECSTaskDefinition_Volume {

	results := map[string]*AWSECSTaskDefinition_Volume{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_Volume{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_VolumeWithName retrieves all AWSECSTaskDefinition_Volume items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_Volume(name string, template *Template) (*AWSECSTaskDefinition_Volume, error) {

	result := &AWSECSTaskDefinition_Volume{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_Volume{}, errors.New("resource not found")

}
