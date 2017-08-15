package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::NetworkInterface AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html
type AWSEC2NetworkInterface struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-description
	Description string `json:"Description"`

	// GroupSet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-groupset
	GroupSet []string `json:"GroupSet"`

	// InterfaceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-interfacetype
	InterfaceType string `json:"InterfaceType"`

	// Ipv6AddressCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-ipv6addresscount
	Ipv6AddressCount int64 `json:"Ipv6AddressCount"`

	// Ipv6Addresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-ipv6addresses
	Ipv6Addresses AWSEC2NetworkInterface_InstanceIpv6Address `json:"Ipv6Addresses"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-privateipaddress
	PrivateIpAddress string `json:"PrivateIpAddress"`

	// PrivateIpAddresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-privateipaddresses
	PrivateIpAddresses []AWSEC2NetworkInterface_PrivateIpAddressSpecification `json:"PrivateIpAddresses"`

	// SecondaryPrivateIpAddressCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-secondaryprivateipcount
	SecondaryPrivateIpAddressCount int64 `json:"SecondaryPrivateIpAddressCount"`

	// SourceDestCheck AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-sourcedestcheck
	SourceDestCheck bool `json:"SourceDestCheck"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-subnetid
	SubnetId string `json:"SubnetId"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkInterface) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterface"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkInterface) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkInterfaceResources retrieves all AWSEC2NetworkInterface items from a CloudFormation template
func GetAllAWSEC2NetworkInterface(template *Template) map[string]*AWSEC2NetworkInterface {

	results := map[string]*AWSEC2NetworkInterface{}
	for name, resource := range template.Resources {
		result := &AWSEC2NetworkInterface{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkInterfaceWithName retrieves all AWSEC2NetworkInterface items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2NetworkInterface(name string, template *Template) (*AWSEC2NetworkInterface, error) {

	result := &AWSEC2NetworkInterface{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkInterface{}, errors.New("resource not found")

}
