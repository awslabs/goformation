package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.NetworkInterface AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html
type AWSEC2Instance_NetworkInterface struct {

	// AssociatePublicIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-associatepubip
	AssociatePublicIpAddress bool `json:"AssociatePublicIpAddress"`

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-delete
	DeleteOnTermination bool `json:"DeleteOnTermination"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-description
	Description string `json:"Description"`

	// DeviceIndex AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-deviceindex
	DeviceIndex string `json:"DeviceIndex"`

	// GroupSet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-groupset
	GroupSet []string `json:"GroupSet"`

	// Ipv6AddressCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#cfn-ec2-instance-networkinterface-ipv6addresscount
	Ipv6AddressCount int64 `json:"Ipv6AddressCount"`

	// Ipv6Addresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#cfn-ec2-instance-networkinterface-ipv6addresses
	Ipv6Addresses []AWSEC2Instance_InstanceIpv6Address `json:"Ipv6Addresses"`

	// NetworkInterfaceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-network-iface
	NetworkInterfaceId string `json:"NetworkInterfaceId"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-privateipaddress
	PrivateIpAddress string `json:"PrivateIpAddress"`

	// PrivateIpAddresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-privateipaddresses
	PrivateIpAddresses []AWSEC2Instance_PrivateIpAddressSpecification `json:"PrivateIpAddresses"`

	// SecondaryPrivateIpAddressCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-secondprivateip
	SecondaryPrivateIpAddressCount int64 `json:"SecondaryPrivateIpAddressCount"`

	// SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-iface-embedded.html#aws-properties-ec2-network-iface-embedded-subnetid
	SubnetId string `json:"SubnetId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_NetworkInterface) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.NetworkInterface"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_NetworkInterface) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_NetworkInterfaceResources retrieves all AWSEC2Instance_NetworkInterface items from a CloudFormation template
func GetAllAWSEC2Instance_NetworkInterface(template *Template) map[string]*AWSEC2Instance_NetworkInterface {

	results := map[string]*AWSEC2Instance_NetworkInterface{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_NetworkInterface{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_NetworkInterfaceWithName retrieves all AWSEC2Instance_NetworkInterface items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_NetworkInterface(name string, template *Template) (*AWSEC2Instance_NetworkInterface, error) {

	result := &AWSEC2Instance_NetworkInterface{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_NetworkInterface{}, errors.New("resource not found")

}
