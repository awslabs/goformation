package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2DHCPOptions AWS CloudFormation Resource (AWS::EC2::DHCPOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html
type AWSEC2DHCPOptions struct {

	// DomainName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-domainname
	DomainName string `json:"DomainName"`

	// DomainNameServers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-domainnameservers
	DomainNameServers []string `json:"DomainNameServers"`

	// NetbiosNameServers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-netbiosnameservers
	NetbiosNameServers []string `json:"NetbiosNameServers"`

	// NetbiosNodeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-netbiosnodetype
	NetbiosNodeType int `json:"NetbiosNodeType"`

	// NtpServers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-ntpservers
	NtpServers []string `json:"NtpServers"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html#cfn-ec2-dhcpoptions-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2DHCPOptions) AWSCloudFormationType() string {
	return "AWS::EC2::DHCPOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2DHCPOptions) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2DHCPOptionsResources retrieves all AWSEC2DHCPOptions items from a CloudFormation template
func (t *Template) GetAllAWSEC2DHCPOptionsResources() map[string]*AWSEC2DHCPOptions {

	results := map[string]*AWSEC2DHCPOptions{}
	for name, resource := range t.Resources {
		result := &AWSEC2DHCPOptions{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2DHCPOptionsWithName retrieves all AWSEC2DHCPOptions items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2DHCPOptionsWithName(name string) (*AWSEC2DHCPOptions, error) {

	result := &AWSEC2DHCPOptions{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2DHCPOptions{}, errors.New("resource not found")

}
