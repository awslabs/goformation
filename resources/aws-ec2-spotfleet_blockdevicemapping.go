package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SpotFleet.BlockDeviceMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html
type AWSEC2SpotFleet_BlockDeviceMapping struct {

	// DeviceName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-devicename

	DeviceName string `json:"DeviceName"`

	// Ebs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-ebs

	Ebs AWSEC2SpotFleet_EbsBlockDevice `json:"Ebs"`

	// NoDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-nodevice

	NoDevice string `json:"NoDevice"`

	// VirtualName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-blockdevicemappings.html#cfn-ec2-spotfleet-blockdevicemapping-virtualname

	VirtualName string `json:"VirtualName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_BlockDeviceMapping) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.BlockDeviceMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_BlockDeviceMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleet_BlockDeviceMappingResources retrieves all AWSEC2SpotFleet_BlockDeviceMapping items from a CloudFormation template
func GetAllAWSEC2SpotFleet_BlockDeviceMapping(template *Template) map[string]*AWSEC2SpotFleet_BlockDeviceMapping {

	results := map[string]*AWSEC2SpotFleet_BlockDeviceMapping{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet_BlockDeviceMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleet_BlockDeviceMappingWithName retrieves all AWSEC2SpotFleet_BlockDeviceMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SpotFleet_BlockDeviceMapping(name string, template *Template) (*AWSEC2SpotFleet_BlockDeviceMapping, error) {

	result := &AWSEC2SpotFleet_BlockDeviceMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet_BlockDeviceMapping{}, errors.New("resource not found")

}
