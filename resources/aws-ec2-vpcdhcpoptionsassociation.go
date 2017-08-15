package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::VPCDHCPOptionsAssociation AWS CloudFormation Resource
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
func GetAllAWSEC2VPCDHCPOptionsAssociation(template *Template) map[string]*AWSEC2VPCDHCPOptionsAssociation {

	results := map[string]*AWSEC2VPCDHCPOptionsAssociation{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPCDHCPOptionsAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCDHCPOptionsAssociationWithName retrieves all AWSEC2VPCDHCPOptionsAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2VPCDHCPOptionsAssociation(name string, template *Template) (*AWSEC2VPCDHCPOptionsAssociation, error) {

	result := &AWSEC2VPCDHCPOptionsAssociation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCDHCPOptionsAssociation{}, errors.New("resource not found")

}
