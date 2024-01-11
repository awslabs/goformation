// Code generated by "go generate". Please don't change this file directly.

package mediatailor

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Channel_TimeShiftConfiguration AWS CloudFormation Resource (AWS::MediaTailor::Channel.TimeShiftConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-timeshiftconfiguration.html
type Channel_TimeShiftConfiguration struct {

	// MaxTimeDelaySeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-timeshiftconfiguration.html#cfn-mediatailor-channel-timeshiftconfiguration-maxtimedelayseconds
	MaxTimeDelaySeconds float64 `json:"MaxTimeDelaySeconds"`

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
func (r *Channel_TimeShiftConfiguration) AWSCloudFormationType() string {
	return "AWS::MediaTailor::Channel.TimeShiftConfiguration"
}
