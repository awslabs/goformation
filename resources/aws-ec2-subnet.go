package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2Subnet AWS CloudFormation Resource (AWS::EC2::Subnet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html
type AWSEC2Subnet struct {

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone"`

	// CidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-cidrblock
	CidrBlock string `json:"CidrBlock"`

	// MapPublicIpOnLaunch AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-mappubliciponlaunch
	MapPublicIpOnLaunch bool `json:"MapPublicIpOnLaunch"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-tags
	Tags []Tag `json:"Tags"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-awsec2subnet-prop-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Subnet) AWSCloudFormationType() string {
	return "AWS::EC2::Subnet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Subnet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SubnetResources retrieves all AWSEC2Subnet items from a CloudFormation template
func (t *Template) GetAllAWSEC2SubnetResources() map[string]*AWSEC2Subnet {

	results := map[string]*AWSEC2Subnet{}
	for name, resource := range t.Resources {
		result := &AWSEC2Subnet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SubnetWithName retrieves all AWSEC2Subnet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetWithName(name string) (*AWSEC2Subnet, error) {

	result := &AWSEC2Subnet{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Subnet{}, errors.New("resource not found")

}
