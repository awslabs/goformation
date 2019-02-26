package cloudformation

// AWSIoTAnalyticsDataset_TriggeringDataset AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.TriggeringDataset)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-triggeringdataset.html
type AWSIoTAnalyticsDataset_TriggeringDataset struct {

	// DatasetName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-triggeringdataset.html#cfn-iotanalytics-dataset-triggeringdataset-datasetname
	DatasetName string `json:"DatasetName,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_TriggeringDataset) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.TriggeringDataset"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_TriggeringDataset) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
