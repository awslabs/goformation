package cloudformation

import (
	"encoding/json"
)

// AWSCodeDeployDeploymentConfig_MinimumHealthyHosts AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html
type AWSCodeDeployDeploymentConfig_MinimumHealthyHosts struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-type
	Type *Value `json:"Type,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-value
	Value *Value `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentConfig_MinimumHealthyHosts) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts"
}

func (r *AWSCodeDeployDeploymentConfig_MinimumHealthyHosts) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
