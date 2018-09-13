package cloudformation

import (
	"encoding/json"
)

// AWSCodeDeployDeploymentGroup_ELBInfo AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.ELBInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-elbinfo.html
type AWSCodeDeployDeploymentGroup_ELBInfo struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-elbinfo.html#cfn-codedeploy-deploymentgroup-elbinfo-name
	Name *Value `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_ELBInfo) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.ELBInfo"
}

func (r *AWSCodeDeployDeploymentGroup_ELBInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
