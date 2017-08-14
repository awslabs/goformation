package resources

// AWS::CodeDeploy::DeploymentGroup.TriggerConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html
type AWSCodeDeployDeploymentGroupTriggerConfig struct {

	// TriggerEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggerevents
	TriggerEvents []AWSCodeDeployDeploymentGroupTriggerConfigstring `json:"TriggerEvents"`

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
func (r *AWSCodeDeployDeploymentGroupTriggerConfig) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.TriggerConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroupTriggerConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
