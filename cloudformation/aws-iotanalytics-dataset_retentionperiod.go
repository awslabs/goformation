package cloudformation

// AWSIoTAnalyticsDataset_RetentionPeriod AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.RetentionPeriod)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-retentionperiod.html
type AWSIoTAnalyticsDataset_RetentionPeriod struct {

	// NumberOfDays AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-retentionperiod.html#cfn-iotanalytics-dataset-retentionperiod-numberofdays
	NumberOfDays int `json:"NumberOfDays,omitempty"`

	// Unlimited AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-retentionperiod.html#cfn-iotanalytics-dataset-retentionperiod-unlimited
	Unlimited bool `json:"Unlimited,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_RetentionPeriod) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.RetentionPeriod"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_RetentionPeriod) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
