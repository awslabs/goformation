package resources

// AWS::IoT::TopicRule.Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html
type AWSIoTTopicRuleAction struct {

	// CloudwatchAlarm AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-cloudwatchalarm
	CloudwatchAlarm AWSIoTTopicRuleActionCloudwatchAlarmAction `json:"CloudwatchAlarm"`

	// CloudwatchMetric AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-cloudwatchmetric
	CloudwatchMetric AWSIoTTopicRuleActionCloudwatchMetricAction `json:"CloudwatchMetric"`

	// DynamoDB AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-dynamodb
	DynamoDB AWSIoTTopicRuleActionDynamoDBAction `json:"DynamoDB"`

	// Elasticsearch AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-elasticsearch
	Elasticsearch AWSIoTTopicRuleActionElasticsearchAction `json:"Elasticsearch"`

	// Firehose AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-firehose
	Firehose AWSIoTTopicRuleActionFirehoseAction `json:"Firehose"`

	// Kinesis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-kinesis
	Kinesis AWSIoTTopicRuleActionKinesisAction `json:"Kinesis"`

	// Lambda AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-lambda
	Lambda AWSIoTTopicRuleActionLambdaAction `json:"Lambda"`

	// Republish AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-republish
	Republish AWSIoTTopicRuleActionRepublishAction `json:"Republish"`

	// S3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-s3
	S3 AWSIoTTopicRuleActionS3Action `json:"S3"`

	// Sns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-sns
	Sns AWSIoTTopicRuleActionSnsAction `json:"Sns"`

	// Sqs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-sqs
	Sqs AWSIoTTopicRuleActionSqsAction `json:"Sqs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRuleAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRuleAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
