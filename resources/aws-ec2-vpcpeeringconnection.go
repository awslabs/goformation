package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::VPCPeeringConnection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
type AWSEC2VPCPeeringConnection struct {

	// PeerOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerownerid
	PeerOwnerId string `json:"PeerOwnerId"`

	// PeerRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerrolearn
	PeerRoleArn string `json:"PeerRoleArn"`

	// PeerVpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peervpcid
	PeerVpcId string `json:"PeerVpcId"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-tags
	Tags []Tag `json:"Tags"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCPeeringConnection) AWSCloudFormationType() string {
	return "AWS::EC2::VPCPeeringConnection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCPeeringConnection) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPCPeeringConnectionResources retrieves all AWSEC2VPCPeeringConnection items from a CloudFormation template
func GetAllAWSEC2VPCPeeringConnectionResources(template *Template) map[string]*AWSEC2VPCPeeringConnection {

	results := map[string]*AWSEC2VPCPeeringConnection{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPCPeeringConnection{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCPeeringConnectionWithName retrieves all AWSEC2VPCPeeringConnection items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2VPCPeeringConnectionWithName(name string, template *Template) (*AWSEC2VPCPeeringConnection, error) {

	result := &AWSEC2VPCPeeringConnection{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCPeeringConnection{}, errors.New("resource not found")

}
