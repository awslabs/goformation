package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::Service.PlacementStrategy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html
type AWSECSService_PlacementStrategy struct {

	// Field AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-field

	Field string `json:"Field"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_PlacementStrategy) AWSCloudFormationType() string {
	return "AWS::ECS::Service.PlacementStrategy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSService_PlacementStrategy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSService_PlacementStrategyResources retrieves all AWSECSService_PlacementStrategy items from a CloudFormation template
func GetAllAWSECSService_PlacementStrategy(template *Template) map[string]*AWSECSService_PlacementStrategy {

	results := map[string]*AWSECSService_PlacementStrategy{}
	for name, resource := range template.Resources {
		result := &AWSECSService_PlacementStrategy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSService_PlacementStrategyWithName retrieves all AWSECSService_PlacementStrategy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSService_PlacementStrategy(name string, template *Template) (*AWSECSService_PlacementStrategy, error) {

	result := &AWSECSService_PlacementStrategy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSService_PlacementStrategy{}, errors.New("resource not found")

}
