package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPC AWS CloudFormation Resource (AWS::EC2::VPC)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html
type AWSEC2VPC struct {

	// CidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-cidrblock
	CidrBlock string `json:"CidrBlock"`

	// EnableDnsHostnames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-EnableDnsHostnames
	EnableDnsHostnames bool `json:"EnableDnsHostnames"`

	// EnableDnsSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-EnableDnsSupport
	EnableDnsSupport bool `json:"EnableDnsSupport"`

	// InstanceTenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-instancetenancy
	InstanceTenancy string `json:"InstanceTenancy"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-aws-ec2-vpc-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPC) AWSCloudFormationType() string {
	return "AWS::EC2::VPC"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPC) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPCResources retrieves all AWSEC2VPC items from a CloudFormation template
func (t *Template) GetAllAWSEC2VPCResources() map[string]*AWSEC2VPC {

	results := map[string]*AWSEC2VPC{}
	for name, resource := range t.Resources {
		result := &AWSEC2VPC{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCWithName retrieves all AWSEC2VPC items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCWithName(name string) (*AWSEC2VPC, error) {

	result := &AWSEC2VPC{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPC{}, errors.New("resource not found")

}
