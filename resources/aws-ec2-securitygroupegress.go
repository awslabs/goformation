package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2SecurityGroupEgress AWS CloudFormation Resource (AWS::EC2::SecurityGroupEgress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html
type AWSEC2SecurityGroupEgress struct {

	// CidrIp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-cidrip
	CidrIp string `json:"CidrIp"`

	// CidrIpv6 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-cidripv6
	CidrIpv6 string `json:"CidrIpv6"`

	// DestinationPrefixListId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-destinationprefixlistid
	DestinationPrefixListId string `json:"DestinationPrefixListId"`

	// DestinationSecurityGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-destinationsecuritygroupid
	DestinationSecurityGroupId string `json:"DestinationSecurityGroupId"`

	// FromPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-fromport
	FromPort int `json:"FromPort"`

	// GroupId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-groupid
	GroupId string `json:"GroupId"`

	// IpProtocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-ipprotocol
	IpProtocol string `json:"IpProtocol"`

	// ToPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-toport
	ToPort int `json:"ToPort"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SecurityGroupEgress) AWSCloudFormationType() string {
	return "AWS::EC2::SecurityGroupEgress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SecurityGroupEgress) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SecurityGroupEgressResources retrieves all AWSEC2SecurityGroupEgress items from a CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupEgressResources() map[string]*AWSEC2SecurityGroupEgress {

	results := map[string]*AWSEC2SecurityGroupEgress{}
	for name, resource := range t.Resources {
		result := &AWSEC2SecurityGroupEgress{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SecurityGroupEgressWithName retrieves all AWSEC2SecurityGroupEgress items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupEgressWithName(name string) (*AWSEC2SecurityGroupEgress, error) {

	result := &AWSEC2SecurityGroupEgress{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SecurityGroupEgress{}, errors.New("resource not found")

}
