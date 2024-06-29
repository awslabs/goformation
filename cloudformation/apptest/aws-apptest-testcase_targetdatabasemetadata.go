// Code generated by "go generate". Please don't change this file directly.

package apptest

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// TestCase_TargetDatabaseMetadata AWS CloudFormation Resource (AWS::AppTest::TestCase.TargetDatabaseMetadata)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-targetdatabasemetadata.html
type TestCase_TargetDatabaseMetadata struct {

	// CaptureTool AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-targetdatabasemetadata.html#cfn-apptest-testcase-targetdatabasemetadata-capturetool
	CaptureTool string `json:"CaptureTool"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-targetdatabasemetadata.html#cfn-apptest-testcase-targetdatabasemetadata-type
	Type string `json:"Type"`

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
func (r *TestCase_TargetDatabaseMetadata) AWSCloudFormationType() string {
	return "AWS::AppTest::TestCase.TargetDatabaseMetadata"
}
