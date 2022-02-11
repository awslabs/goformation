package kafkaconnect

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Connector_ScaleOutPolicy AWS CloudFormation Resource (AWS::KafkaConnect::Connector.ScaleOutPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-scaleoutpolicy.html
type Connector_ScaleOutPolicy struct {

	// CpuUtilizationPercentage AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-scaleoutpolicy.html#cfn-kafkaconnect-connector-scaleoutpolicy-cpuutilizationpercentage
	CpuUtilizationPercentage int `json:"CpuUtilizationPercentage"`

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
func (r *Connector_ScaleOutPolicy) AWSCloudFormationType() string {
	return "AWS::KafkaConnect::Connector.ScaleOutPolicy"
}
