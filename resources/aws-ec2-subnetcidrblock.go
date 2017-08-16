package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2SubnetCidrBlock AWS CloudFormation Resource (AWS::EC2::SubnetCidrBlock)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html
type AWSEC2SubnetCidrBlock struct {

	// Ipv6CidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-ipv6cidrblock
	Ipv6CidrBlock string `json:"Ipv6CidrBlock"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-subnetid
	SubnetId string `json:"SubnetId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SubnetCidrBlock) AWSCloudFormationType() string {
	return "AWS::EC2::SubnetCidrBlock"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SubnetCidrBlock) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SubnetCidrBlockResources retrieves all AWSEC2SubnetCidrBlock items from a CloudFormation template
func (t *Template) GetAllAWSEC2SubnetCidrBlockResources() map[string]*AWSEC2SubnetCidrBlock {

	results := map[string]*AWSEC2SubnetCidrBlock{}
	for name, resource := range t.Resources {
		result := &AWSEC2SubnetCidrBlock{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SubnetCidrBlockWithName retrieves all AWSEC2SubnetCidrBlock items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetCidrBlockWithName(name string) (*AWSEC2SubnetCidrBlock, error) {

	result := &AWSEC2SubnetCidrBlock{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SubnetCidrBlock{}, errors.New("resource not found")

}
