package cloudformation

import (
	"encoding/json"
)

// AWSAutoScalingLaunchConfiguration_BlockDeviceMapping AWS CloudFormation Resource (AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html
type AWSAutoScalingLaunchConfiguration_BlockDeviceMapping struct {

	// DeviceName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-devicename
	DeviceName *Value `json:"DeviceName,omitempty"`

	// Ebs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-ebs
	Ebs *AWSAutoScalingLaunchConfiguration_BlockDevice `json:"Ebs,omitempty"`

	// NoDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-nodevice
	NoDevice *Value `json:"NoDevice,omitempty"`

	// VirtualName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-virtualname
	VirtualName *Value `json:"VirtualName,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingLaunchConfiguration_BlockDeviceMapping) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping"
}

func (r *AWSAutoScalingLaunchConfiguration_BlockDeviceMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
