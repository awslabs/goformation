package serverless

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Function_EventInvokeDestinationConfig AWS CloudFormation Resource (AWS::Serverless::Function.EventInvokeDestinationConfig)
// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#event-invoke-destination-config-object
type Function_EventInvokeDestinationConfig struct {

	// OnFailure AWS CloudFormation Property
	// Required: true
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#event-invoke-destination-config-object
	OnFailure *Function_Destination `json:"OnFailure"`

	// OnSuccess AWS CloudFormation Property
	// Required: true
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#event-invoke-destination-config-object
	OnSuccess *Function_Destination `json:"OnSuccess"`

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
func (r *Function_EventInvokeDestinationConfig) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.EventInvokeDestinationConfig"
}
