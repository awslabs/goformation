package ec2

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// LaunchTemplate_CapacityReservationPreference AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.CapacityReservationPreference)
// See:
type LaunchTemplate_CapacityReservationPreference struct {

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *LaunchTemplate_CapacityReservationPreference) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.CapacityReservationPreference"
}
