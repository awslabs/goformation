package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPCDHCPOptionsAssociation AWS CloudFormation Resource (AWS::EC2::VPCDHCPOptionsAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html
type AWSEC2VPCDHCPOptionsAssociation struct {

	// DhcpOptionsId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html#cfn-ec2-vpcdhcpoptionsassociation-dhcpoptionsid
	DhcpOptionsId string `json:"DhcpOptionsId"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html#cfn-ec2-vpcdhcpoptionsassociation-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCDHCPOptionsAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::VPCDHCPOptionsAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCDHCPOptionsAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPCDHCPOptionsAssociationResources retrieves all AWSEC2VPCDHCPOptionsAssociation items from a CloudFormation template
func (t *Template) GetAllAWSEC2VPCDHCPOptionsAssociationResources() map[string]*AWSEC2VPCDHCPOptionsAssociation {

	results := map[string]*AWSEC2VPCDHCPOptionsAssociation{}
	for name, resource := range t.Resources {
		result := &AWSEC2VPCDHCPOptionsAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCDHCPOptionsAssociationWithName retrieves all AWSEC2VPCDHCPOptionsAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCDHCPOptionsAssociationWithName(name string) (*AWSEC2VPCDHCPOptionsAssociation, error) {

	result := &AWSEC2VPCDHCPOptionsAssociation{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCDHCPOptionsAssociation{}, errors.New("resource not found")

}
