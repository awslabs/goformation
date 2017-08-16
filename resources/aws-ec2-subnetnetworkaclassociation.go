package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2SubnetNetworkAclAssociation AWS CloudFormation Resource (AWS::EC2::SubnetNetworkAclAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html
type AWSEC2SubnetNetworkAclAssociation struct {

	// NetworkAclId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-networkaclid
	NetworkAclId string `json:"NetworkAclId"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-associationid
	SubnetId string `json:"SubnetId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SubnetNetworkAclAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::SubnetNetworkAclAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SubnetNetworkAclAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SubnetNetworkAclAssociationResources retrieves all AWSEC2SubnetNetworkAclAssociation items from a CloudFormation template
func (t *Template) GetAllAWSEC2SubnetNetworkAclAssociationResources() map[string]*AWSEC2SubnetNetworkAclAssociation {

	results := map[string]*AWSEC2SubnetNetworkAclAssociation{}
	for name, resource := range t.Resources {
		result := &AWSEC2SubnetNetworkAclAssociation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SubnetNetworkAclAssociationWithName retrieves all AWSEC2SubnetNetworkAclAssociation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetNetworkAclAssociationWithName(name string) (*AWSEC2SubnetNetworkAclAssociation, error) {

	result := &AWSEC2SubnetNetworkAclAssociation{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SubnetNetworkAclAssociation{}, errors.New("resource not found")

}
