package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.TriggerConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html
type AWSCodeDeployDeploymentGroup_TriggerConfig struct {

	// TriggerEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggerevents
	TriggerEvents []string `json:"TriggerEvents"`

	// TriggerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggername
	TriggerName string `json:"TriggerName"`

	// TriggerTargetArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggertargetarn
	TriggerTargetArn string `json:"TriggerTargetArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_TriggerConfig) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.TriggerConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_TriggerConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
