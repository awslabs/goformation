// Code generated by "go generate". Please don't change this file directly.

package mediaconnect

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Flow_MediaStreamAttributes AWS CloudFormation Resource (AWS::MediaConnect::Flow.MediaStreamAttributes)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamattributes.html
type Flow_MediaStreamAttributes struct {

	// Fmtp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamattributes.html#cfn-mediaconnect-flow-mediastreamattributes-fmtp
	Fmtp *Flow_Fmtp `json:"Fmtp,omitempty"`

	// Lang AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastreamattributes.html#cfn-mediaconnect-flow-mediastreamattributes-lang
	Lang *string `json:"Lang,omitempty"`

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
func (r *Flow_MediaStreamAttributes) AWSCloudFormationType() string {
	return "AWS::MediaConnect::Flow.MediaStreamAttributes"
}
