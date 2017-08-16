package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSAutoScalingLifecycleHook AWS CloudFormation Resource (AWS::AutoScaling::LifecycleHook)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html
type AWSAutoScalingLifecycleHook struct {

	// AutoScalingGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-autoscalinggroupname
	AutoScalingGroupName string `json:"AutoScalingGroupName"`

	// DefaultResult AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-defaultresult
	DefaultResult string `json:"DefaultResult"`

	// HeartbeatTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-heartbeattimeout
	HeartbeatTimeout int `json:"HeartbeatTimeout"`

	// LifecycleTransition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-lifecycletransition
	LifecycleTransition string `json:"LifecycleTransition"`

	// NotificationMetadata AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-notificationmetadata
	NotificationMetadata string `json:"NotificationMetadata"`

	// NotificationTargetARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-notificationtargetarn
	NotificationTargetARN string `json:"NotificationTargetARN"`

	// RoleARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-rolearn
	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingLifecycleHook) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LifecycleHook"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingLifecycleHook) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingLifecycleHookResources retrieves all AWSAutoScalingLifecycleHook items from a CloudFormation template
func (t *Template) GetAllAWSAutoScalingLifecycleHookResources() map[string]*AWSAutoScalingLifecycleHook {

	results := map[string]*AWSAutoScalingLifecycleHook{}
	for name, resource := range t.Resources {
		result := &AWSAutoScalingLifecycleHook{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingLifecycleHookWithName retrieves all AWSAutoScalingLifecycleHook items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingLifecycleHookWithName(name string) (*AWSAutoScalingLifecycleHook, error) {

	result := &AWSAutoScalingLifecycleHook{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingLifecycleHook{}, errors.New("resource not found")

}
