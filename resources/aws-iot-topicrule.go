package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
type AWSIoTTopicRule struct {

	// RuleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-rulename
	RuleName string `json:"RuleName"`

	// TopicRulePayload AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-topicrulepayload
	TopicRulePayload AWSIoTTopicRule_TopicRulePayload `json:"TopicRulePayload"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRuleResources retrieves all AWSIoTTopicRule items from a CloudFormation template
func GetAllAWSIoTTopicRule(template *Template) map[string]*AWSIoTTopicRule {

	results := map[string]*AWSIoTTopicRule{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRuleWithName retrieves all AWSIoTTopicRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule(name string, template *Template) (*AWSIoTTopicRule, error) {

	result := &AWSIoTTopicRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule{}, errors.New("resource not found")

}
