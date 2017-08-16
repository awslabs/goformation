package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2NetworkInterfacePermission AWS CloudFormation Resource (AWS::EC2::NetworkInterfacePermission)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html
type AWSEC2NetworkInterfacePermission struct {

	// AwsAccountId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html#cfn-ec2-networkinterfacepermission-awsaccountid
	AwsAccountId string `json:"AwsAccountId"`

	// NetworkInterfaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html#cfn-ec2-networkinterfacepermission-networkinterfaceid
	NetworkInterfaceId string `json:"NetworkInterfaceId"`

	// Permission AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html#cfn-ec2-networkinterfacepermission-permission
	Permission string `json:"Permission"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkInterfacePermission) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterfacePermission"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkInterfacePermission) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkInterfacePermissionResources retrieves all AWSEC2NetworkInterfacePermission items from a CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfacePermissionResources() map[string]*AWSEC2NetworkInterfacePermission {

	results := map[string]*AWSEC2NetworkInterfacePermission{}
	for name, resource := range t.Resources {
		result := &AWSEC2NetworkInterfacePermission{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkInterfacePermissionWithName retrieves all AWSEC2NetworkInterfacePermission items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfacePermissionWithName(name string) (*AWSEC2NetworkInterfacePermission, error) {

	result := &AWSEC2NetworkInterfacePermission{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkInterfacePermission{}, errors.New("resource not found")

}
