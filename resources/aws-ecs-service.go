package resources

// AWS::ECS::Service AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
type AWSECSService struct {

	// Cluster AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-cluster

	Cluster string `json:"Cluster"`

	// DeploymentConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-deploymentconfiguration

	DeploymentConfiguration AWSECSService_DeploymentConfiguration `json:"DeploymentConfiguration"`

	// DesiredCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-desiredcount

	DesiredCount int64 `json:"DesiredCount"`

	// LoadBalancers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-loadbalancers

	LoadBalancers []AWSECSService_LoadBalancer `json:"LoadBalancers"`

	// PlacementConstraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementconstraints

	PlacementConstraints []AWSECSService_PlacementConstraint `json:"PlacementConstraints"`

	// PlacementStrategies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-placementstrategies

	PlacementStrategies []AWSECSService_PlacementStrategy `json:"PlacementStrategies"`

	// Role AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-role

	Role string `json:"Role"`

	// ServiceName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-servicename

	ServiceName string `json:"ServiceName"`

	// TaskDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-taskdefinition

	TaskDefinition string `json:"TaskDefinition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService) AWSCloudFormationType() string {
	return "AWS::ECS::Service"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSService) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
