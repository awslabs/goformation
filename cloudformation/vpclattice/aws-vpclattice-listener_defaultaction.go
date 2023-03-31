// Code generated by "go generate". Please don't change this file directly.

package vpclattice

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Listener_DefaultAction AWS CloudFormation Resource (AWS::VpcLattice::Listener.DefaultAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html
type Listener_DefaultAction struct {

	// FixedResponse AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html#cfn-vpclattice-listener-defaultaction-fixedresponse
	FixedResponse *Listener_FixedResponse `json:"FixedResponse,omitempty"`

	// Forward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-listener-defaultaction.html#cfn-vpclattice-listener-defaultaction-forward
	Forward *Listener_Forward `json:"Forward,omitempty"`

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
func (r *Listener_DefaultAction) AWSCloudFormationType() string {
	return "AWS::VpcLattice::Listener.DefaultAction"
}
