package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSStepFunctionsStateMachine AWS CloudFormation Resource (AWS::StepFunctions::StateMachine)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
type AWSStepFunctionsStateMachine struct {

	// DefinitionString AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-definitionstring
	DefinitionString string `json:"DefinitionString"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html#cfn-stepfunctions-statemachine-rolearn
	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSStepFunctionsStateMachine) AWSCloudFormationType() string {
	return "AWS::StepFunctions::StateMachine"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSStepFunctionsStateMachine) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSStepFunctionsStateMachineResources retrieves all AWSStepFunctionsStateMachine items from a CloudFormation template
func (t *Template) GetAllAWSStepFunctionsStateMachineResources() map[string]*AWSStepFunctionsStateMachine {

	results := map[string]*AWSStepFunctionsStateMachine{}
	for name, resource := range t.Resources {
		result := &AWSStepFunctionsStateMachine{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSStepFunctionsStateMachineWithName retrieves all AWSStepFunctionsStateMachine items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsStateMachineWithName(name string) (*AWSStepFunctionsStateMachine, error) {

	result := &AWSStepFunctionsStateMachine{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSStepFunctionsStateMachine{}, errors.New("resource not found")

}
