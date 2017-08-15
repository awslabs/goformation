package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::AutoScaling::AutoScalingGroup.NotificationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html
type AWSAutoScalingAutoScalingGroup_NotificationConfiguration struct {

	// NotificationTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html#cfn-as-group-notificationconfigurations-notificationtypes
	NotificationTypes []string `json:"NotificationTypes"`

	// TopicARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html#cfn-autoscaling-autoscalinggroup-notificationconfigurations-topicarn
	TopicARN string `json:"TopicARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_NotificationConfiguration) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.NotificationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroup_NotificationConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingAutoScalingGroup_NotificationConfigurationResources retrieves all AWSAutoScalingAutoScalingGroup_NotificationConfiguration items from a CloudFormation template
func GetAllAWSAutoScalingAutoScalingGroup_NotificationConfiguration(template *Template) map[string]*AWSAutoScalingAutoScalingGroup_NotificationConfiguration {

	results := map[string]*AWSAutoScalingAutoScalingGroup_NotificationConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSAutoScalingAutoScalingGroup_NotificationConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingAutoScalingGroup_NotificationConfigurationWithName retrieves all AWSAutoScalingAutoScalingGroup_NotificationConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSAutoScalingAutoScalingGroup_NotificationConfiguration(name string, template *Template) (*AWSAutoScalingAutoScalingGroup_NotificationConfiguration, error) {

	result := &AWSAutoScalingAutoScalingGroup_NotificationConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingAutoScalingGroup_NotificationConfiguration{}, errors.New("resource not found")

}
