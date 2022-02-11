package kafkaconnect

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Connector_WorkerLogDelivery AWS CloudFormation Resource (AWS::KafkaConnect::Connector.WorkerLogDelivery)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html
type Connector_WorkerLogDelivery struct {

	// CloudWatchLogs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html#cfn-kafkaconnect-connector-workerlogdelivery-cloudwatchlogs
	CloudWatchLogs *Connector_CloudWatchLogsLogDelivery `json:"CloudWatchLogs,omitempty"`

	// Firehose AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html#cfn-kafkaconnect-connector-workerlogdelivery-firehose
	Firehose *Connector_FirehoseLogDelivery `json:"Firehose,omitempty"`

	// S3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html#cfn-kafkaconnect-connector-workerlogdelivery-s3
	S3 *Connector_S3LogDelivery `json:"S3,omitempty"`

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
func (r *Connector_WorkerLogDelivery) AWSCloudFormationType() string {
	return "AWS::KafkaConnect::Connector.WorkerLogDelivery"
}
