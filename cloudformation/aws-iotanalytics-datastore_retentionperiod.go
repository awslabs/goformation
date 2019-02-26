package cloudformation

// AWSIoTAnalyticsDatastore_RetentionPeriod AWS CloudFormation Resource (AWS::IoTAnalytics::Datastore.RetentionPeriod)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html
type AWSIoTAnalyticsDatastore_RetentionPeriod struct {

	// NumberOfDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html#cfn-iotanalytics-datastore-retentionperiod-numberofdays
	NumberOfDays int `json:"NumberOfDays,omitempty"`

	// Unlimited AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html#cfn-iotanalytics-datastore-retentionperiod-unlimited
	Unlimited bool `json:"Unlimited,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDatastore_RetentionPeriod) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Datastore.RetentionPeriod"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDatastore_RetentionPeriod) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
