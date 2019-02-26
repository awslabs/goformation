package cloudformation

// AWSIoTAnalyticsChannel_RetentionPeriod AWS CloudFormation Resource (AWS::IoTAnalytics::Channel.RetentionPeriod)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-retentionperiod.html
type AWSIoTAnalyticsChannel_RetentionPeriod struct {

	// NumberOfDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-retentionperiod.html#cfn-iotanalytics-channel-retentionperiod-numberofdays
	NumberOfDays int `json:"NumberOfDays,omitempty"`

	// Unlimited AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-channel-retentionperiod.html#cfn-iotanalytics-channel-retentionperiod-unlimited
	Unlimited bool `json:"Unlimited,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsChannel_RetentionPeriod) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Channel.RetentionPeriod"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsChannel_RetentionPeriod) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
