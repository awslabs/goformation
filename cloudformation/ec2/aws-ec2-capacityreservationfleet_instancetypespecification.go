package ec2

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// CapacityReservationFleet_InstanceTypeSpecification AWS CloudFormation Resource (AWS::EC2::CapacityReservationFleet.InstanceTypeSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html
type CapacityReservationFleet_InstanceTypeSpecification struct {

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`

	// AvailabilityZoneId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-availabilityzoneid
	AvailabilityZoneId string `json:"AvailabilityZoneId,omitempty"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized,omitempty"`

	// InstancePlatform AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-instanceplatform
	InstancePlatform string `json:"InstancePlatform,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// Priority AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-priority
	Priority int `json:"Priority,omitempty"`

	// Weight AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-capacityreservationfleet-instancetypespecification.html#cfn-ec2-capacityreservationfleet-instancetypespecification-weight
	Weight float64 `json:"Weight,omitempty"`

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
func (r *CapacityReservationFleet_InstanceTypeSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::CapacityReservationFleet.InstanceTypeSpecification"
}
