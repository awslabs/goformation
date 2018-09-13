package cloudformation

import (
	"encoding/json"
)

// AWSECSTaskDefinition_TaskDefinitionPlacementConstraint AWS CloudFormation Resource (AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html
type AWSECSTaskDefinition_TaskDefinitionPlacementConstraint struct {

	// Expression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-expression
	Expression *Value `json:"Expression,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-taskdefinitionplacementconstraint.html#cfn-ecs-taskdefinition-taskdefinitionplacementconstraint-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_TaskDefinitionPlacementConstraint) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.TaskDefinitionPlacementConstraint"
}

func (r *AWSECSTaskDefinition_TaskDefinitionPlacementConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
