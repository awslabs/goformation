package cloudformation

// AWSIoTAnalyticsDataset_Schedule AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.Schedule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-trigger-schedule.html
type AWSIoTAnalyticsDataset_Schedule struct {

	// ScheduleExpression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-trigger-schedule.html#cfn-iotanalytics-dataset-trigger-schedule-scheduleexpression
	ScheduleExpression string `json:"ScheduleExpression,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_Schedule) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.Schedule"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_Schedule) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
