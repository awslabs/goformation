package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.SqsAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html
type AWSIoTTopicRule_SqsAction struct {

	// QueueUrl AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-queueurl
	QueueUrl string `json:"QueueUrl"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-rolearn
	RoleArn string `json:"RoleArn"`

	// UseBase64 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sqs.html#cfn-iot-sqs-usebase64
	UseBase64 bool `json:"UseBase64"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_SqsAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.SqsAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_SqsAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_SqsActionResources retrieves all AWSIoTTopicRule_SqsAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_SqsAction(template *Template) map[string]*AWSIoTTopicRule_SqsAction {

	results := map[string]*AWSIoTTopicRule_SqsAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_SqsAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_SqsActionWithName retrieves all AWSIoTTopicRule_SqsAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_SqsAction(name string, template *Template) (*AWSIoTTopicRule_SqsAction, error) {

	result := &AWSIoTTopicRule_SqsAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_SqsAction{}, errors.New("resource not found")

}
