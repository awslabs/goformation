package iotanalytics

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Dataset_Schedule AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.Schedule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-schedule.html
type Dataset_Schedule struct {

	// ScheduleExpression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-schedule.html#cfn-iotanalytics-dataset-schedule-scheduleexpression
	ScheduleExpression string `json:"ScheduleExpression"`

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
func (r *Dataset_Schedule) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.Schedule"
}
