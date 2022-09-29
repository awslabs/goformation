// Code generated by "go generate". Please don't change this file directly.

package lakeformation

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// PrincipalPermissions_TableWildcard AWS CloudFormation Resource (AWS::LakeFormation::PrincipalPermissions.TableWildcard)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewildcard.html
type PrincipalPermissions_TableWildcard struct {

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
func (r *PrincipalPermissions_TableWildcard) AWSCloudFormationType() string {
	return "AWS::LakeFormation::PrincipalPermissions.TableWildcard"
}
