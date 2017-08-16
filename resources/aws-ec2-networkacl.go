package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2NetworkAcl AWS CloudFormation Resource (AWS::EC2::NetworkAcl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
type AWSEC2NetworkAcl struct {

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-tags
	Tags []Tag `json:"Tags"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAcl) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAcl"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkAcl) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkAclResources retrieves all AWSEC2NetworkAcl items from a CloudFormation template
func (t *Template) GetAllAWSEC2NetworkAclResources() map[string]*AWSEC2NetworkAcl {

	results := map[string]*AWSEC2NetworkAcl{}
	for name, resource := range t.Resources {
		result := &AWSEC2NetworkAcl{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkAclWithName retrieves all AWSEC2NetworkAcl items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkAclWithName(name string) (*AWSEC2NetworkAcl, error) {

	result := &AWSEC2NetworkAcl{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkAcl{}, errors.New("resource not found")

}
