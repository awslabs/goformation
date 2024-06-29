// Code generated by "go generate". Please don't change this file directly.

package apptest

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// TestCase_InputFile AWS CloudFormation Resource (AWS::AppTest::TestCase.InputFile)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html
type TestCase_InputFile struct {

	// FileMetadata AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-filemetadata
	FileMetadata *TestCase_FileMetadata `json:"FileMetadata"`

	// SourceLocation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-sourcelocation
	SourceLocation string `json:"SourceLocation"`

	// TargetLocation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-inputfile.html#cfn-apptest-testcase-inputfile-targetlocation
	TargetLocation string `json:"TargetLocation"`

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
func (r *TestCase_InputFile) AWSCloudFormationType() string {
	return "AWS::AppTest::TestCase.InputFile"
}