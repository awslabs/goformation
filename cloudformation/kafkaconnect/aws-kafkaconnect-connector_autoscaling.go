package kafkaconnect

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Connector_AutoScaling AWS CloudFormation Resource (AWS::KafkaConnect::Connector.AutoScaling)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html
type Connector_AutoScaling struct {

	// MaxWorkerCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-maxworkercount
	MaxWorkerCount int `json:"MaxWorkerCount"`

	// McuCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-mcucount
	McuCount int `json:"McuCount"`

	// MinWorkerCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-minworkercount
	MinWorkerCount int `json:"MinWorkerCount"`

	// ScaleInPolicy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-scaleinpolicy
	ScaleInPolicy *Connector_ScaleInPolicy `json:"ScaleInPolicy"`

	// ScaleOutPolicy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-autoscaling.html#cfn-kafkaconnect-connector-autoscaling-scaleoutpolicy
	ScaleOutPolicy *Connector_ScaleOutPolicy `json:"ScaleOutPolicy"`

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
func (r *Connector_AutoScaling) AWSCloudFormationType() string {
	return "AWS::KafkaConnect::Connector.AutoScaling"
}
