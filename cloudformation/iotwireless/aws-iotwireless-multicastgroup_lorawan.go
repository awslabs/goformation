package iotwireless

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// MulticastGroup_LoRaWAN AWS CloudFormation Resource (AWS::IoTWireless::MulticastGroup.LoRaWAN)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html
type MulticastGroup_LoRaWAN struct {

	// DlClass AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-dlclass
	DlClass string `json:"DlClass"`

	// NumberOfDevicesInGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-numberofdevicesingroup
	NumberOfDevicesInGroup *int `json:"NumberOfDevicesInGroup,omitempty"`

	// NumberOfDevicesRequested AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-numberofdevicesrequested
	NumberOfDevicesRequested *int `json:"NumberOfDevicesRequested,omitempty"`

	// RfRegion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-rfregion
	RfRegion string `json:"RfRegion"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *MulticastGroup_LoRaWAN) AWSCloudFormationType() string {
	return "AWS::IoTWireless::MulticastGroup.LoRaWAN"
}
