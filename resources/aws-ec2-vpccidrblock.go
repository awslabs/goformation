package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::VPCCidrBlock AWS CloudFormation Resource
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
func GetAllAWSEC2VPCCidrBlock(template *Template) map[string]*AWSEC2VPCCidrBlock {

	results := map[string]*AWSEC2VPCCidrBlock{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPCCidrBlock{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCCidrBlockWithName retrieves all AWSEC2VPCCidrBlock items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2VPCCidrBlock(name string, template *Template) (*AWSEC2VPCCidrBlock, error) {

	result := &AWSEC2VPCCidrBlock{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCCidrBlock{}, errors.New("resource not found")

}
