package resources

// AWS::ECS::Cluster AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
type AWSECSCluster struct {

	// ClusterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustername
	ClusterName string `json:"ClusterName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSCluster) AWSCloudFormationType() string {
	return "AWS::ECS::Cluster"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECSCluster) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
