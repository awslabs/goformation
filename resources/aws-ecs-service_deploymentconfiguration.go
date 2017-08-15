package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::Service.DeploymentConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
type AWSECSService_DeploymentConfiguration struct {

	// MaximumPercent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-maximumpercent

	MaximumPercent int64 `json:"MaximumPercent"`

	// MinimumHealthyPercent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-minimumhealthypercent

	MinimumHealthyPercent int64 `json:"MinimumHealthyPercent"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_DeploymentConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::Service.DeploymentConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSService_DeploymentConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSService_DeploymentConfigurationResources retrieves all AWSECSService_DeploymentConfiguration items from a CloudFormation template
func GetAllAWSECSService_DeploymentConfiguration(template *Template) map[string]*AWSECSService_DeploymentConfiguration {

	results := map[string]*AWSECSService_DeploymentConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSECSService_DeploymentConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSService_DeploymentConfigurationWithName retrieves all AWSECSService_DeploymentConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSService_DeploymentConfiguration(name string, template *Template) (*AWSECSService_DeploymentConfiguration, error) {

	result := &AWSECSService_DeploymentConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSService_DeploymentConfiguration{}, errors.New("resource not found")

}
