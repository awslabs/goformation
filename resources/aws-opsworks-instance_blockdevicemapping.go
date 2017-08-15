package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Instance.BlockDeviceMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html
type AWSOpsWorksInstance_BlockDeviceMapping struct {

	// DeviceName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-devicename
	DeviceName string `json:"DeviceName"`

	// Ebs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-ebs
	Ebs AWSOpsWorksInstance_EbsBlockDevice `json:"Ebs"`

	// NoDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-nodevice
	NoDevice string `json:"NoDevice"`

	// VirtualName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-virtualname
	VirtualName string `json:"VirtualName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksInstance_BlockDeviceMapping) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Instance.BlockDeviceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksInstance_BlockDeviceMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
