package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.Ulimit AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html
type AWSECSTaskDefinition_Ulimit struct {

	// HardLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-hardlimit

	HardLimit int64 `json:"HardLimit"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-name

	Name string `json:"Name"`

	// SoftLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html#cfn-ecs-taskdefinition-containerdefinition-ulimit-softlimit

	SoftLimit int64 `json:"SoftLimit"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_Ulimit) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.Ulimit"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_Ulimit) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_UlimitResources retrieves all AWSECSTaskDefinition_Ulimit items from a CloudFormation template
func GetAllAWSECSTaskDefinition_Ulimit(template *Template) map[string]*AWSECSTaskDefinition_Ulimit {

	results := map[string]*AWSECSTaskDefinition_Ulimit{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_Ulimit{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_UlimitWithName retrieves all AWSECSTaskDefinition_Ulimit items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_Ulimit(name string, template *Template) (*AWSECSTaskDefinition_Ulimit, error) {

	result := &AWSECSTaskDefinition_Ulimit{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_Ulimit{}, errors.New("resource not found")

}
