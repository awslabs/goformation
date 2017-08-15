package resources

// AWSECSService_DeploymentConfiguration AWS CloudFormation Resource (AWS::ECS::Service.DeploymentConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
type AWSECSService_DeploymentConfiguration struct {

	// MaximumPercent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-maximumpercent
	MaximumPercent int64 `json:"MaximumPercent"`

	// MinimumHealthyPercent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-minimumhealthypercent
	MinimumHealthyPercent int64 `json:"MinimumHealthyPercent"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_DeploymentConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::Service.DeploymentConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSService_DeploymentConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
