package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.FirehoseAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html
type AWSIoTTopicRule_FirehoseAction struct {

	// DeliveryStreamName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html#cfn-iot-firehose-deliverystreamname

	DeliveryStreamName string `json:"DeliveryStreamName"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html#cfn-iot-firehose-rolearn

	RoleArn string `json:"RoleArn"`

	// Separator AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-firehose.html#cfn-iot-firehose-separator

	Separator string `json:"Separator"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_FirehoseAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.FirehoseAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_FirehoseAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_FirehoseActionResources retrieves all AWSIoTTopicRule_FirehoseAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_FirehoseAction(template *Template) map[string]*AWSIoTTopicRule_FirehoseAction {

	results := map[string]*AWSIoTTopicRule_FirehoseAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_FirehoseAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_FirehoseActionWithName retrieves all AWSIoTTopicRule_FirehoseAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_FirehoseAction(name string, template *Template) (*AWSIoTTopicRule_FirehoseAction, error) {

	result := &AWSIoTTopicRule_FirehoseAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_FirehoseAction{}, errors.New("resource not found")

}
