package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.NoDevice AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-nodevice.html
type AWSEC2Instance_NoDevice struct {
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_NoDevice) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.NoDevice"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_NoDevice) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_NoDeviceResources retrieves all AWSEC2Instance_NoDevice items from a CloudFormation template
func GetAllAWSEC2Instance_NoDevice(template *Template) map[string]*AWSEC2Instance_NoDevice {

	results := map[string]*AWSEC2Instance_NoDevice{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_NoDevice{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_NoDeviceWithName retrieves all AWSEC2Instance_NoDevice items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_NoDevice(name string, template *Template) (*AWSEC2Instance_NoDevice, error) {

	result := &AWSEC2Instance_NoDevice{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_NoDevice{}, errors.New("resource not found")

}
