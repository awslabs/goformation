package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.SnsAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html
type AWSIoTTopicRule_SnsAction struct {

	// MessageFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-snsaction
	MessageFormat string `json:"MessageFormat"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-rolearn
	RoleArn string `json:"RoleArn"`

	// TargetArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-sns.html#cfn-iot-sns-targetarn
	TargetArn string `json:"TargetArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_SnsAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.SnsAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_SnsAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_SnsActionResources retrieves all AWSIoTTopicRule_SnsAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_SnsAction(template *Template) map[string]*AWSIoTTopicRule_SnsAction {

	results := map[string]*AWSIoTTopicRule_SnsAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_SnsAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_SnsActionWithName retrieves all AWSIoTTopicRule_SnsAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_SnsAction(name string, template *Template) (*AWSIoTTopicRule_SnsAction, error) {

	result := &AWSIoTTopicRule_SnsAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_SnsAction{}, errors.New("resource not found")

}
