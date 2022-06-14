// Code generated by "go generate". Please don't change this file directly.

package networkmanager

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// CoreNetwork_CoreNetworkSegment AWS CloudFormation Resource (AWS::NetworkManager::CoreNetwork.CoreNetworkSegment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html
type CoreNetwork_CoreNetworkSegment struct {

	// EdgeLocations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html#cfn-networkmanager-corenetwork-corenetworksegment-edgelocations
	EdgeLocations *[]string `json:"EdgeLocations,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html#cfn-networkmanager-corenetwork-corenetworksegment-name
	Name *string `json:"Name,omitempty"`

	// SharedSegments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-corenetwork-corenetworksegment.html#cfn-networkmanager-corenetwork-corenetworksegment-sharedsegments
	SharedSegments *[]string `json:"SharedSegments,omitempty"`

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
func (r *CoreNetwork_CoreNetworkSegment) AWSCloudFormationType() string {
	return "AWS::NetworkManager::CoreNetwork.CoreNetworkSegment"
}
