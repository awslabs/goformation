package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Host AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
type AWSEC2Host struct {

	// AutoPlacement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement
	AutoPlacement string `json:"AutoPlacement"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype
	InstanceType string `json:"InstanceType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Host) AWSCloudFormationType() string {
	return "AWS::EC2::Host"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Host) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2HostResources retrieves all AWSEC2Host items from a CloudFormation template
func GetAllAWSEC2HostResources(template *Template) map[string]*AWSEC2Host {

	results := map[string]*AWSEC2Host{}
	for name, resource := range template.Resources {
		result := &AWSEC2Host{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2HostWithName retrieves all AWSEC2Host items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2HostWithName(name string, template *Template) (*AWSEC2Host, error) {

	result := &AWSEC2Host{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Host{}, errors.New("resource not found")

}
