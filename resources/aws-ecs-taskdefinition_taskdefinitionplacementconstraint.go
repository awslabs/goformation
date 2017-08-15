package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html
type AWSECSTaskDefinition_TaskDefinitionPlacementConstraint struct {

	// Expression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-expression
	Expression string `json:"Expression"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_TaskDefinitionPlacementConstraint) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSTaskDefinition_TaskDefinitionPlacementConstraint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSECSTaskDefinition_TaskDefinitionPlacementConstraintResources retrieves all AWSECSTaskDefinition_TaskDefinitionPlacementConstraint items from a CloudFormation template
func GetAllAWSECSTaskDefinition_TaskDefinitionPlacementConstraint(template *Template) map[string]*AWSECSTaskDefinition_TaskDefinitionPlacementConstraint {

	results := map[string]*AWSECSTaskDefinition_TaskDefinitionPlacementConstraint{}
	for name, resource := range template.Resources {
		result := &AWSECSTaskDefinition_TaskDefinitionPlacementConstraint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSECSTaskDefinition_TaskDefinitionPlacementConstraintWithName retrieves all AWSECSTaskDefinition_TaskDefinitionPlacementConstraint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSECSTaskDefinition_TaskDefinitionPlacementConstraint(name string, template *Template) (*AWSECSTaskDefinition_TaskDefinitionPlacementConstraint, error) {

	result := &AWSECSTaskDefinition_TaskDefinitionPlacementConstraint{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSECSTaskDefinition_TaskDefinitionPlacementConstraint{}, errors.New("resource not found")

}
