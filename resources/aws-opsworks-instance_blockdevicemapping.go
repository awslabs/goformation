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

// GetAllAWSOpsWorksInstance_BlockDeviceMappingResources retrieves all AWSOpsWorksInstance_BlockDeviceMapping items from a CloudFormation template
func GetAllAWSOpsWorksInstance_BlockDeviceMapping(template *Template) map[string]*AWSOpsWorksInstance_BlockDeviceMapping {

	results := map[string]*AWSOpsWorksInstance_BlockDeviceMapping{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksInstance_BlockDeviceMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksInstance_BlockDeviceMappingWithName retrieves all AWSOpsWorksInstance_BlockDeviceMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksInstance_BlockDeviceMapping(name string, template *Template) (*AWSOpsWorksInstance_BlockDeviceMapping, error) {

	result := &AWSOpsWorksInstance_BlockDeviceMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksInstance_BlockDeviceMapping{}, errors.New("resource not found")

}
