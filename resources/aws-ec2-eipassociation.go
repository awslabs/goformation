package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2EIPAssociation AWS CloudFormation Resource (AWS::EC2::EIPAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
type AWSEC2EIPAssociation struct {

	// AllocationId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-allocationid
	AllocationId string `json:"AllocationId"`

	// EIP AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-eip
	EIP string `json:"EIP"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-instanceid
	InstanceId string `json:"InstanceId"`

	// NetworkInterfaceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-networkinterfaceid
	NetworkInterfaceId string `json:"NetworkInterfaceId"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html#cfn-ec2-eipassociation-PrivateIpAddress
	PrivateIpAddress string `json:"PrivateIpAddress"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EIPAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::EIPAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2EIPAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2EIPAssociationResources retrieves all AWSEC2EIPAssociation items from a CloudFormation template
func (t *Template) GetAllAWSEC2EIPAssociationResources() map[string]*AWSEC2EIPAssociation {

	results := map[string]*AWSEC2EIPAssociation{}
	for name, resource := range t.Resources {
		result := &AWSEC2EIPAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2EIPAssociationWithName retrieves all AWSEC2EIPAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EIPAssociationWithName(name string) (*AWSEC2EIPAssociation, error) {

	result := &AWSEC2EIPAssociation{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2EIPAssociation{}, errors.New("resource not found")

}
