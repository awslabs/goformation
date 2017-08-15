package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.TopicRulePayload AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html
type AWSIoTTopicRule_TopicRulePayload struct {

	// Actions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-actions

	Actions []AWSIoTTopicRule_Action `json:"Actions"`

	// AwsIotSqlVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-awsiotsqlversion

	AwsIotSqlVersion string `json:"AwsIotSqlVersion"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-description

	Description string `json:"Description"`

	// RuleDisabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-ruledisabled

	RuleDisabled bool `json:"RuleDisabled"`

	// Sql AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrulepayload.html#cfn-iot-topicrulepayload-sql

	Sql string `json:"Sql"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_TopicRulePayload) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.TopicRulePayload"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_TopicRulePayload) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_TopicRulePayloadResources retrieves all AWSIoTTopicRule_TopicRulePayload items from a CloudFormation template
func GetAllAWSIoTTopicRule_TopicRulePayload(template *Template) map[string]*AWSIoTTopicRule_TopicRulePayload {

	results := map[string]*AWSIoTTopicRule_TopicRulePayload{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_TopicRulePayload{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_TopicRulePayloadWithName retrieves all AWSIoTTopicRule_TopicRulePayload items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_TopicRulePayload(name string, template *Template) (*AWSIoTTopicRule_TopicRulePayload, error) {

	result := &AWSIoTTopicRule_TopicRulePayload{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_TopicRulePayload{}, errors.New("resource not found")

}
