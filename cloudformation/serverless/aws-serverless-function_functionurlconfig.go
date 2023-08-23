// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Function_FunctionUrlConfig AWS CloudFormation Resource (AWS::Serverless::Function.FunctionUrlConfig)
// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-functionurlconfig.html
type Function_FunctionUrlConfig struct {

	// AuthType AWS CloudFormation Property
	// Required: true
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-functionurlconfig.html#sam-function-functionurlconfig-authtype
	AuthType string `json:"AuthType"`

	// Cors AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-functionurlconfig.html#sam-function-functionurlconfig-cors
	Cors *Function_Cors `json:"Cors,omitempty"`

	// InvokeMode AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-functionurlconfig.html#sam-function-functionurlconfig-invokemode
	InvokeMode *string `json:"InvokeMode,omitempty"`

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
func (r *Function_FunctionUrlConfig) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.FunctionUrlConfig"
}
