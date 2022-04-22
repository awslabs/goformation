// Code generated by "go generate". Please don't change this file directly.

package iottwinmaker

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// ComponentType_Relationship AWS CloudFormation Resource (AWS::IoTTwinMaker::ComponentType.Relationship)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationship.html
type ComponentType_Relationship struct {

	// RelationshipType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationship.html#cfn-iottwinmaker-componenttype-relationship-relationshiptype
	RelationshipType *string `json:"RelationshipType,omitempty"`

	// TargetComponentTypeId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationship.html#cfn-iottwinmaker-componenttype-relationship-targetcomponenttypeid
	TargetComponentTypeId *string `json:"TargetComponentTypeId,omitempty"`

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
func (r *ComponentType_Relationship) AWSCloudFormationType() string {
	return "AWS::IoTTwinMaker::ComponentType.Relationship"
}
