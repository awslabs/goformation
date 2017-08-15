package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.HostEntry AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html
type AWSECSTaskDefinition_HostEntry struct {

	// Hostname AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-hostname
	Hostname string `json:"Hostname"`

	// IpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-hostentry.html#cfn-ecs-taskdefinition-containerdefinition-hostentry-ipaddress
	IpAddress string `json:"IpAddress"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_HostEntry) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.HostEntry"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_HostEntry) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_HostEntryResources retrieves all AWSECSTaskDefinition_HostEntry items from a CloudFormation template
func GetAllAWSECSTaskDefinition_HostEntry(template *Template) map[string]*AWSECSTaskDefinition_HostEntry {

	results := map[string]*AWSECSTaskDefinition_HostEntry{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_HostEntry{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_HostEntryWithName retrieves all AWSECSTaskDefinition_HostEntry items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_HostEntry(name string, template *Template) (*AWSECSTaskDefinition_HostEntry, error) {

	result := &AWSECSTaskDefinition_HostEntry{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_HostEntry{}, errors.New("resource not found")

}
