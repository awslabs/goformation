// Code generated by "go generate". Please don't change this file directly.

package iotevents

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// AlarmModel_AlarmCapabilities AWS CloudFormation Resource (AWS::IoTEvents::AlarmModel.AlarmCapabilities)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmcapabilities.html
type AlarmModel_AlarmCapabilities struct {

	// AcknowledgeFlow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmcapabilities.html#cfn-iotevents-alarmmodel-alarmcapabilities-acknowledgeflow
	AcknowledgeFlow *AlarmModel_AcknowledgeFlow `json:"AcknowledgeFlow,omitempty"`

	// InitializationConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmcapabilities.html#cfn-iotevents-alarmmodel-alarmcapabilities-initializationconfiguration
	InitializationConfiguration *AlarmModel_InitializationConfiguration `json:"InitializationConfiguration,omitempty"`

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
func (r *AlarmModel_AlarmCapabilities) AWSCloudFormationType() string {
	return "AWS::IoTEvents::AlarmModel.AlarmCapabilities"
}