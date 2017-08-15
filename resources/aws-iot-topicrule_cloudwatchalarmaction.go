package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.CloudwatchAlarmAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html
type AWSIoTTopicRule_CloudwatchAlarmAction struct {

	// AlarmName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-alarmname
	AlarmName string `json:"AlarmName"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-rolearn
	RoleArn string `json:"RoleArn"`

	// StateReason AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-statereason
	StateReason string `json:"StateReason"`

	// StateValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchalarm.html#cfn-iot-cloudwatchalarm-statevalue
	StateValue string `json:"StateValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_CloudwatchAlarmAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.CloudwatchAlarmAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_CloudwatchAlarmAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_CloudwatchAlarmActionResources retrieves all AWSIoTTopicRule_CloudwatchAlarmAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_CloudwatchAlarmAction(template *Template) map[string]*AWSIoTTopicRule_CloudwatchAlarmAction {

	results := map[string]*AWSIoTTopicRule_CloudwatchAlarmAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_CloudwatchAlarmAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_CloudwatchAlarmActionWithName retrieves all AWSIoTTopicRule_CloudwatchAlarmAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_CloudwatchAlarmAction(name string, template *Template) (*AWSIoTTopicRule_CloudwatchAlarmAction, error) {

	result := &AWSIoTTopicRule_CloudwatchAlarmAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_CloudwatchAlarmAction{}, errors.New("resource not found")

}
