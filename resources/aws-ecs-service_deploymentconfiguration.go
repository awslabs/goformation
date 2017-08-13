package resources

// AWS::ECS::Service.DeploymentConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
type AWSECSServiceDeploymentConfiguration struct {
    
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
func (r *AWSECSServiceDeploymentConfiguration) AWSCloudFormationType() string {
    return "AWS::ECS::Service.DeploymentConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSServiceDeploymentConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}