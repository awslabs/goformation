package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.RepublishAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-republish.html
type AWSIoTTopicRule_RepublishAction struct {

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-republish.html#cfn-iot-republish-rolearn

	RoleArn string `json:"RoleArn"`

	// Topic AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-republish.html#cfn-iot-republish-topic

	Topic string `json:"Topic"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_RepublishAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.RepublishAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_RepublishAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_RepublishActionResources retrieves all AWSIoTTopicRule_RepublishAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_RepublishAction(template *Template) map[string]*AWSIoTTopicRule_RepublishAction {

	results := map[string]*AWSIoTTopicRule_RepublishAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_RepublishAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_RepublishActionWithName retrieves all AWSIoTTopicRule_RepublishAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_RepublishAction(name string, template *Template) (*AWSIoTTopicRule_RepublishAction, error) {

	result := &AWSIoTTopicRule_RepublishAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_RepublishAction{}, errors.New("resource not found")

}
