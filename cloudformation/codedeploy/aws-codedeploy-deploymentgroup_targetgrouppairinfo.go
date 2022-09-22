// Code generated by "go generate". Please don't change this file directly.

package codedeploy

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// DeploymentGroup_TargetGroupPairInfo AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.TargetGroupPairInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html
type DeploymentGroup_TargetGroupPairInfo struct {

	// ProdTrafficRoute AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html#cfn-codedeploy-deploymentgroup-targetgrouppairinfo-prodtrafficroute
	ProdTrafficRoute *DeploymentGroup_TrafficRoute `json:"ProdTrafficRoute,omitempty"`

	// TargetGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html#cfn-codedeploy-deploymentgroup-targetgrouppairinfo-targetgroups
	TargetGroups []DeploymentGroup_TargetGroupInfo `json:"TargetGroups,omitempty"`

	// TestTrafficRoute AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html#cfn-codedeploy-deploymentgroup-targetgrouppairinfo-testtrafficroute
	TestTrafficRoute *DeploymentGroup_TrafficRoute `json:"TestTrafficRoute,omitempty"`

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
func (r *DeploymentGroup_TargetGroupPairInfo) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.TargetGroupPairInfo"
}
