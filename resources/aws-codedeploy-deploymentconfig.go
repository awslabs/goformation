package resources

// AWS::CodeDeploy::DeploymentConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
type AWSCodeDeployDeploymentConfig struct {

	// DeploymentConfigName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-deploymentconfigname

	DeploymentConfigName string `json:"DeploymentConfigName"`

	// MinimumHealthyHosts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts

	MinimumHealthyHosts AWSCodeDeployDeploymentConfig_MinimumHealthyHosts `json:"MinimumHealthyHosts"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentConfig) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
