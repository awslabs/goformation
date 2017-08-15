package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html
type AWSAutoScalingLaunchConfiguration_BlockDeviceMapping struct {

	// DeviceName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-devicename
	DeviceName string `json:"DeviceName"`

	// Ebs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-ebs
	Ebs AWSAutoScalingLaunchConfiguration_BlockDevice `json:"Ebs"`

	// NoDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-nodevice
	NoDevice bool `json:"NoDevice"`

	// VirtualName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html#cfn-as-launchconfig-blockdev-mapping-virtualname
	VirtualName string `json:"VirtualName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingLaunchConfiguration_BlockDeviceMapping) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LaunchConfiguration.BlockDeviceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingLaunchConfiguration_BlockDeviceMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingLaunchConfiguration_BlockDeviceMappingResources retrieves all AWSAutoScalingLaunchConfiguration_BlockDeviceMapping items from a CloudFormation template
func GetAllAWSAutoScalingLaunchConfiguration_BlockDeviceMapping(template *Template) map[string]*AWSAutoScalingLaunchConfiguration_BlockDeviceMapping {

	results := map[string]*AWSAutoScalingLaunchConfiguration_BlockDeviceMapping{}
	for name, resource := range template.Resources {
		result := &AWSAutoScalingLaunchConfiguration_BlockDeviceMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingLaunchConfiguration_BlockDeviceMappingWithName retrieves all AWSAutoScalingLaunchConfiguration_BlockDeviceMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSAutoScalingLaunchConfiguration_BlockDeviceMapping(name string, template *Template) (*AWSAutoScalingLaunchConfiguration_BlockDeviceMapping, error) {

	result := &AWSAutoScalingLaunchConfiguration_BlockDeviceMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingLaunchConfiguration_BlockDeviceMapping{}, errors.New("resource not found")

}
