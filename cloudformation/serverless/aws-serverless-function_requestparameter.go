// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Function_RequestParameter AWS CloudFormation Resource (AWS::Serverless::Function.RequestParameter)
// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-requestparameter.html
type Function_RequestParameter struct {

	// Caching AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-requestparameter.html#sam-function-requestparameter-caching
	Caching *bool `json:"Caching,omitempty"`

	// Required AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-requestparameter.html#sam-function-requestparameter-required
	Required *bool `json:"Required,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_RequestParameter) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.RequestParameter"
}
