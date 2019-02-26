package cloudformation

// AWSIoTAnalyticsDataset_Action AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html
type AWSIoTAnalyticsDataset_Action struct {

	// ActionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html#cfn-iotanalytics-dataset-action-actionname
	ActionName string `json:"ActionName,omitempty"`

	// ContainerAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html#cfn-iotanalytics-dataset-action-containeraction
	ContainerAction *AWSIoTAnalyticsDataset_ContainerAction `json:"ContainerAction,omitempty"`

	// QueryAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-action.html#cfn-iotanalytics-dataset-action-queryaction
	QueryAction *AWSIoTAnalyticsDataset_QueryAction `json:"QueryAction,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_Action) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.Action"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_Action) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
