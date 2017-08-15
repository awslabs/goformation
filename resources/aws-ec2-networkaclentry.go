package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2NetworkAclEntry AWS CloudFormation Resource (AWS::EC2::NetworkAclEntry)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html
type AWSEC2NetworkAclEntry struct {

	// CidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-cidrblock
	CidrBlock string `json:"CidrBlock"`

	// Egress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-egress
	Egress bool `json:"Egress"`

	// Icmp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-icmp
	Icmp AWSEC2NetworkAclEntry_Icmp `json:"Icmp"`

	// Ipv6CidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-ipv6cidrblock
	Ipv6CidrBlock string `json:"Ipv6CidrBlock"`

	// NetworkAclId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-networkaclid
	NetworkAclId string `json:"NetworkAclId"`

	// PortRange AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-portrange
	PortRange AWSEC2NetworkAclEntry_PortRange `json:"PortRange"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-protocol
	Protocol int64 `json:"Protocol"`

	// RuleAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-ruleaction
	RuleAction string `json:"RuleAction"`

	// RuleNumber AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-rulenumber
	RuleNumber int64 `json:"RuleNumber"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAclEntry) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAclEntry"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkAclEntry) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkAclEntryResources retrieves all AWSEC2NetworkAclEntry items from a CloudFormation template
func GetAllAWSEC2NetworkAclEntryResources(template *Template) map[string]*AWSEC2NetworkAclEntry {

	results := map[string]*AWSEC2NetworkAclEntry{}
	for name, resource := range template.Resources {
		result := &AWSEC2NetworkAclEntry{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkAclEntryWithName retrieves all AWSEC2NetworkAclEntry items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2NetworkAclEntryWithName(name string, template *Template) (*AWSEC2NetworkAclEntry, error) {

	result := &AWSEC2NetworkAclEntry{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkAclEntry{}, errors.New("resource not found")

}
