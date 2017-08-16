package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEFSMountTarget AWS CloudFormation Resource (AWS::EFS::MountTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html
type AWSEFSMountTarget struct {

	// FileSystemId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-filesystemid
	FileSystemId string `json:"FileSystemId"`

	// IpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipaddress
	IpAddress string `json:"IpAddress"`

	// SecurityGroups AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-securitygroups
	SecurityGroups []string `json:"SecurityGroups"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-subnetid
	SubnetId string `json:"SubnetId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEFSMountTarget) AWSCloudFormationType() string {
	return "AWS::EFS::MountTarget"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEFSMountTarget) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEFSMountTargetResources retrieves all AWSEFSMountTarget items from a CloudFormation template
func (t *Template) GetAllAWSEFSMountTargetResources() map[string]*AWSEFSMountTarget {

	results := map[string]*AWSEFSMountTarget{}
	for name, resource := range t.Resources {
		result := &AWSEFSMountTarget{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEFSMountTargetWithName retrieves all AWSEFSMountTarget items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEFSMountTargetWithName(name string) (*AWSEFSMountTarget, error) {

	result := &AWSEFSMountTarget{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEFSMountTarget{}, errors.New("resource not found")

}
