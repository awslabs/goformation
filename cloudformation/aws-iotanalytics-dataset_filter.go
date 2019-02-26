package cloudformation

// AWSIoTAnalyticsDataset_Filter AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.Filter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-filter.html
type AWSIoTAnalyticsDataset_Filter struct {

	// DeltaTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-filter.html#cfn-iotanalytics-dataset-filter-deltatime
	DeltaTime *AWSIoTAnalyticsDataset_DeltaTime `json:"DeltaTime,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_Filter) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.Filter"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_Filter) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
