package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::Service.PlacementConstraint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html
type AWSECSService_PlacementConstraint struct {

	// Expression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-expression

	Expression string `json:"Expression"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementconstraint.html#cfn-ecs-service-placementconstraint-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_PlacementConstraint) AWSCloudFormationType() string {
	return "AWS::ECS::Service.PlacementConstraint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSService_PlacementConstraint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSService_PlacementConstraintResources retrieves all AWSECSService_PlacementConstraint items from a CloudFormation template
func GetAllAWSECSService_PlacementConstraint(template *Template) map[string]*AWSECSService_PlacementConstraint {

	results := map[string]*AWSECSService_PlacementConstraint{}
	for name, resource := range template.Resources {
		result := &AWSECSService_PlacementConstraint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSService_PlacementConstraintWithName retrieves all AWSECSService_PlacementConstraint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSService_PlacementConstraint(name string, template *Template) (*AWSECSService_PlacementConstraint, error) {

	result := &AWSECSService_PlacementConstraint{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSService_PlacementConstraint{}, errors.New("resource not found")

}
