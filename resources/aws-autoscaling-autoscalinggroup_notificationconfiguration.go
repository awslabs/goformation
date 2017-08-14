package resources

// AWS::AutoScaling::AutoScalingGroup.NotificationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-notificationconfigurations.html
type AWSAutoScalingAutoScalingGroupNotificationConfiguration struct {

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
func (r *AWSAutoScalingAutoScalingGroupNotificationConfiguration) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.NotificationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroupNotificationConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
