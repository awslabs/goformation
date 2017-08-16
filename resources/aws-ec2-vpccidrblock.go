package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPCCidrBlock AWS CloudFormation Resource (AWS::EC2::VPCCidrBlock)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html
type AWSEC2VPCCidrBlock struct {

	// AmazonProvidedIpv6CidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-amazonprovidedipv6cidrblock
	AmazonProvidedIpv6CidrBlock bool `json:"AmazonProvidedIpv6CidrBlock"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCCidrBlock) AWSCloudFormationType() string {
	return "AWS::EC2::VPCCidrBlock"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCCidrBlock) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPCCidrBlockResources retrieves all AWSEC2VPCCidrBlock items from a CloudFormation template
func (t *Template) GetAllAWSEC2VPCCidrBlockResources() map[string]*AWSEC2VPCCidrBlock {

	results := map[string]*AWSEC2VPCCidrBlock{}
	for name, resource := range t.Resources {
		result := &AWSEC2VPCCidrBlock{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCCidrBlockWithName retrieves all AWSEC2VPCCidrBlock items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCCidrBlockWithName(name string) (*AWSEC2VPCCidrBlock, error) {

	result := &AWSEC2VPCCidrBlock{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCCidrBlock{}, errors.New("resource not found")

}
