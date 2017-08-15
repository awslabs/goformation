package resources

// AWS::Config::ConfigurationRecorder.RecordingGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html
type AWSConfigConfigurationRecorder_RecordingGroup struct {

	// AllSupported AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-allsupported
	AllSupported bool `json:"AllSupported"`

	// IncludeGlobalResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-includeglobalresourcetypes
	IncludeGlobalResourceTypes bool `json:"IncludeGlobalResourceTypes"`

	// ResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-resourcetypes
	ResourceTypes []string `json:"ResourceTypes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigurationRecorder_RecordingGroup) AWSCloudFormationType() string {
	return "AWS::Config::ConfigurationRecorder.RecordingGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigurationRecorder_RecordingGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
