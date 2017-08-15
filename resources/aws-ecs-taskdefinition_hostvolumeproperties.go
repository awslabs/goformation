package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.HostVolumeProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html
type AWSECSTaskDefinition_HostVolumeProperties struct {

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html#cfn-ecs-taskdefinition-volumes-host-sourcepath
	SourcePath string `json:"SourcePath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_HostVolumeProperties) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.HostVolumeProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_HostVolumeProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_HostVolumePropertiesResources retrieves all AWSECSTaskDefinition_HostVolumeProperties items from a CloudFormation template
func GetAllAWSECSTaskDefinition_HostVolumeProperties(template *Template) map[string]*AWSECSTaskDefinition_HostVolumeProperties {

	results := map[string]*AWSECSTaskDefinition_HostVolumeProperties{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_HostVolumeProperties{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_HostVolumePropertiesWithName retrieves all AWSECSTaskDefinition_HostVolumeProperties items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_HostVolumeProperties(name string, template *Template) (*AWSECSTaskDefinition_HostVolumeProperties, error) {

	result := &AWSECSTaskDefinition_HostVolumeProperties{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_HostVolumeProperties{}, errors.New("resource not found")

}
