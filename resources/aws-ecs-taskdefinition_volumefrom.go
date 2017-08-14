package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.VolumeFrom AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html
type AWSECSTaskDefinition_VolumeFrom struct {

	// ReadOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html#cfn-ecs-taskdefinition-containerdefinition-volumesfrom-readonly

	ReadOnly bool `json:"ReadOnly"`

	// SourceContainer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-volumesfrom.html#cfn-ecs-taskdefinition-containerdefinition-volumesfrom-sourcecontainer

	SourceContainer string `json:"SourceContainer"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_VolumeFrom) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.VolumeFrom"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_VolumeFrom) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_VolumeFromResources retrieves all AWSECSTaskDefinition_VolumeFrom items from a CloudFormation template
func GetAllAWSECSTaskDefinition_VolumeFrom(template *Template) map[string]*AWSECSTaskDefinition_VolumeFrom {

	results := map[string]*AWSECSTaskDefinition_VolumeFrom{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_VolumeFrom{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_VolumeFromWithName retrieves all AWSECSTaskDefinition_VolumeFrom items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_VolumeFrom(name string, template *Template) (*AWSECSTaskDefinition_VolumeFrom, error) {

	result := &AWSECSTaskDefinition_VolumeFrom{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_VolumeFrom{}, errors.New("resource not found")

}
