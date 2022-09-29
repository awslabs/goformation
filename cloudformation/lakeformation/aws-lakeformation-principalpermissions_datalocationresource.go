// Code generated by "go generate". Please don't change this file directly.

package lakeformation

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// PrincipalPermissions_DataLocationResource AWS CloudFormation Resource (AWS::LakeFormation::PrincipalPermissions.DataLocationResource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalocationresource.html
type PrincipalPermissions_DataLocationResource struct {

	// CatalogId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalocationresource.html#cfn-lakeformation-principalpermissions-datalocationresource-catalogid
	CatalogId string `json:"CatalogId"`

	// ResourceArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-datalocationresource.html#cfn-lakeformation-principalpermissions-datalocationresource-resourcearn
	ResourceArn string `json:"ResourceArn"`

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
func (r *PrincipalPermissions_DataLocationResource) AWSCloudFormationType() string {
	return "AWS::LakeFormation::PrincipalPermissions.DataLocationResource"
}
