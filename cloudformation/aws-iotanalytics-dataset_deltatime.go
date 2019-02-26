package cloudformation

// AWSIoTAnalyticsDataset_DeltaTime AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.DeltaTime)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatime.html
type AWSIoTAnalyticsDataset_DeltaTime struct {

	// OffsetSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatime.html#cfn-iotanalytics-dataset-deltatime-offsetseconds
	OffsetSeconds int `json:"OffsetSeconds,omitempty"`

	// TimeExpression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatime.html#cfn-iotanalytics-dataset-deltatime-timeexpression
	TimeExpression string `json:"TimeExpression,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_DeltaTime) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.DeltaTime"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_DeltaTime) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
