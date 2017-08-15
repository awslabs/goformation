package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.KinesisAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html
type AWSIoTTopicRule_KinesisAction struct {

	// PartitionKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html#cfn-iot-kinesis-partitionkey
	PartitionKey string `json:"PartitionKey"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html#cfn-iot-kinesis-rolearn
	RoleArn string `json:"RoleArn"`

	// StreamName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-kinesis.html#cfn-iot-kinesis-streamname
	StreamName string `json:"StreamName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_KinesisAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.KinesisAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_KinesisAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_KinesisActionResources retrieves all AWSIoTTopicRule_KinesisAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_KinesisAction(template *Template) map[string]*AWSIoTTopicRule_KinesisAction {

	results := map[string]*AWSIoTTopicRule_KinesisAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_KinesisAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_KinesisActionWithName retrieves all AWSIoTTopicRule_KinesisAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_KinesisAction(name string, template *Template) (*AWSIoTTopicRule_KinesisAction, error) {

	result := &AWSIoTTopicRule_KinesisAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_KinesisAction{}, errors.New("resource not found")

}
