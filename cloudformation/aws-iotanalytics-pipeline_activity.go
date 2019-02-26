package cloudformation

// AWSIoTAnalyticsPipeline_Activity AWS CloudFormation Resource (AWS::IoTAnalytics::Pipeline.Activity)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html
type AWSIoTAnalyticsPipeline_Activity struct {

	// AddAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-addattributes
	AddAttributes *AWSIoTAnalyticsPipeline_AddAttributes `json:"AddAttributes,omitempty"`

	// Channel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-channel
	Channel *AWSIoTAnalyticsPipeline_Channel `json:"Channel,omitempty"`

	// Datastore AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-datastore
	Datastore *AWSIoTAnalyticsPipeline_Datastore `json:"Datastore,omitempty"`

	// DeviceRegistryEnrich AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-deviceregistryenrich
	DeviceRegistryEnrich *AWSIoTAnalyticsPipeline_DeviceRegistryEnrich `json:"DeviceRegistryEnrich,omitempty"`

	// DeviceShadowEnrich AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-deviceshadowenrich
	DeviceShadowEnrich *AWSIoTAnalyticsPipeline_DeviceShadowEnrich `json:"DeviceShadowEnrich,omitempty"`

	// Filter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-filter
	Filter *AWSIoTAnalyticsPipeline_Filter `json:"Filter,omitempty"`

	// Lambda AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-lambda
	Lambda *AWSIoTAnalyticsPipeline_Lambda `json:"Lambda,omitempty"`

	// Math AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-math
	Math *AWSIoTAnalyticsPipeline_Math `json:"Math,omitempty"`

	// RemoveAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-removeattributes
	RemoveAttributes *AWSIoTAnalyticsPipeline_RemoveAttributes `json:"RemoveAttributes,omitempty"`

	// SelectAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-activity.html#cfn-iotanalytics-pipeline-activity-selectattributes
	SelectAttributes *AWSIoTAnalyticsPipeline_SelectAttributes `json:"SelectAttributes,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTAnalyticsPipeline_Activity) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Pipeline.Activity"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSIoTAnalyticsPipeline_Activity) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
