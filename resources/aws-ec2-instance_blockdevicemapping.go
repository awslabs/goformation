package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.BlockDeviceMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
type AWSEC2Instance_BlockDeviceMapping struct {

	// DeviceName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-devicename
	DeviceName string `json:"DeviceName"`

	// Ebs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-ebs
	Ebs AWSEC2Instance_Ebs `json:"Ebs"`

	// NoDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-nodevice
	NoDevice AWSEC2Instance_NoDevice `json:"NoDevice"`

	// VirtualName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html#cfn-ec2-blockdev-mapping-virtualname
	VirtualName string `json:"VirtualName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_BlockDeviceMapping) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.BlockDeviceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_BlockDeviceMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_BlockDeviceMappingResources retrieves all AWSEC2Instance_BlockDeviceMapping items from a CloudFormation template
func GetAllAWSEC2Instance_BlockDeviceMapping(template *Template) map[string]*AWSEC2Instance_BlockDeviceMapping {

	results := map[string]*AWSEC2Instance_BlockDeviceMapping{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_BlockDeviceMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_BlockDeviceMappingWithName retrieves all AWSEC2Instance_BlockDeviceMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_BlockDeviceMapping(name string, template *Template) (*AWSEC2Instance_BlockDeviceMapping, error) {

	result := &AWSEC2Instance_BlockDeviceMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_BlockDeviceMapping{}, errors.New("resource not found")

}
