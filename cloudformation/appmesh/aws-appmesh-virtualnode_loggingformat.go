// Code generated by "go generate". Please don't change this file directly.

package appmesh

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// VirtualNode_LoggingFormat AWS CloudFormation Resource (AWS::AppMesh::VirtualNode.LoggingFormat)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-loggingformat.html
type VirtualNode_LoggingFormat struct {

	// Json AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-loggingformat.html#cfn-appmesh-virtualnode-loggingformat-json
	Json []VirtualNode_JsonFormatRef `json:"Json,omitempty"`

	// Text AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-loggingformat.html#cfn-appmesh-virtualnode-loggingformat-text
	Text *string `json:"Text,omitempty"`

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
func (r *VirtualNode_LoggingFormat) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualNode.LoggingFormat"
}
