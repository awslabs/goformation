// Code generated by "go generate". Please don't change this file directly.

package mediaconvert

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobTemplate_HopDestination AWS CloudFormation Resource (AWS::MediaConvert::JobTemplate.HopDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html
type JobTemplate_HopDestination struct {

	// Priority AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-priority
	Priority *int `json:"Priority,omitempty"`

	// Queue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-queue
	Queue *string `json:"Queue,omitempty"`

	// WaitMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconvert-jobtemplate-hopdestination.html#cfn-mediaconvert-jobtemplate-hopdestination-waitminutes
	WaitMinutes *int `json:"WaitMinutes,omitempty"`

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
func (r *JobTemplate_HopDestination) AWSCloudFormationType() string {
	return "AWS::MediaConvert::JobTemplate.HopDestination"
}
