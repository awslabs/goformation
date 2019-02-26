package cloudformation

// AWSIoTAnalyticsDataset_QueryAction AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.QueryAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html
type AWSIoTAnalyticsDataset_QueryAction struct {

	// Filters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-filters
	Filters []AWSIoTAnalyticsDataset_Filter `json:"Filters,omitempty"`

	// SqlQuery AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-queryaction.html#cfn-iotanalytics-dataset-queryaction-sqlquery
	SqlQuery string `json:"SqlQuery,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsDataset_QueryAction) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.QueryAction"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsDataset_QueryAction) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
