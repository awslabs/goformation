package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html
type AWSIoTTopicRule_Action struct {

	// CloudwatchAlarm AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-cloudwatchalarm
	CloudwatchAlarm AWSIoTTopicRule_CloudwatchAlarmAction `json:"CloudwatchAlarm"`

	// CloudwatchMetric AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-cloudwatchmetric
	CloudwatchMetric AWSIoTTopicRule_CloudwatchMetricAction `json:"CloudwatchMetric"`

	// DynamoDB AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-dynamodb
	DynamoDB AWSIoTTopicRule_DynamoDBAction `json:"DynamoDB"`

	// Elasticsearch AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-elasticsearch
	Elasticsearch AWSIoTTopicRule_ElasticsearchAction `json:"Elasticsearch"`

	// Firehose AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-firehose
	Firehose AWSIoTTopicRule_FirehoseAction `json:"Firehose"`

	// Kinesis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-kinesis
	Kinesis AWSIoTTopicRule_KinesisAction `json:"Kinesis"`

	// Lambda AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-lambda
	Lambda AWSIoTTopicRule_LambdaAction `json:"Lambda"`

	// Republish AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-republish
	Republish AWSIoTTopicRule_RepublishAction `json:"Republish"`

	// S3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-s3
	S3 AWSIoTTopicRule_S3Action `json:"S3"`

	// Sns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-sns
	Sns AWSIoTTopicRule_SnsAction `json:"Sns"`

	// Sqs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-actions.html#cfn-iot-action-sqs
	Sqs AWSIoTTopicRule_SqsAction `json:"Sqs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_Action) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_Action) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_ActionResources retrieves all AWSIoTTopicRule_Action items from a CloudFormation template
func GetAllAWSIoTTopicRule_Action(template *Template) map[string]*AWSIoTTopicRule_Action {

	results := map[string]*AWSIoTTopicRule_Action{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_Action{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_ActionWithName retrieves all AWSIoTTopicRule_Action items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_Action(name string, template *Template) (*AWSIoTTopicRule_Action, error) {

	result := &AWSIoTTopicRule_Action{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_Action{}, errors.New("resource not found")

}
