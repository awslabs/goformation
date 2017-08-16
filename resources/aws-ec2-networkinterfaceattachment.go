package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2NetworkInterfaceAttachment AWS CloudFormation Resource (AWS::EC2::NetworkInterfaceAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html
type AWSEC2NetworkInterfaceAttachment struct {

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-deleteonterm
	DeleteOnTermination bool `json:"DeleteOnTermination"`

	// DeviceIndex AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-deviceindex
	DeviceIndex string `json:"DeviceIndex"`

	// InstanceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-instanceid
	InstanceId string `json:"InstanceId"`

	// NetworkInterfaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html#cfn-ec2-network-interface-attachment-networkinterfaceid
	NetworkInterfaceId string `json:"NetworkInterfaceId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkInterfaceAttachment) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterfaceAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkInterfaceAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkInterfaceAttachmentResources retrieves all AWSEC2NetworkInterfaceAttachment items from a CloudFormation template
func (t *Template) GetAllAWSEC2NetworkInterfaceAttachmentResources() map[string]*AWSEC2NetworkInterfaceAttachment {

	results := map[string]*AWSEC2NetworkInterfaceAttachment{}
	for name, resource := range t.Resources {
		result := &AWSEC2NetworkInterfaceAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkInterfaceAttachmentWithName retrieves all AWSEC2NetworkInterfaceAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkInterfaceAttachmentWithName(name string) (*AWSEC2NetworkInterfaceAttachment, error) {

	result := &AWSEC2NetworkInterfaceAttachment{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkInterfaceAttachment{}, errors.New("resource not found")

}
