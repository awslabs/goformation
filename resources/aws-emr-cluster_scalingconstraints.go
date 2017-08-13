package resources

// AWS::EMR::Cluster.ScalingConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html
type AWSEMRClusterScalingConstraints struct {
    
    // MaxCapacity AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html#cfn-elasticmapreduce-cluster-scalingconstraints-maxcapacity
    MaxCapacity int64 `json:"MaxCapacity"`
    
    // MinCapacity AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html#cfn-elasticmapreduce-cluster-scalingconstraints-mincapacity
    MinCapacity int64 `json:"MinCapacity"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterScalingConstraints) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.ScalingConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterScalingConstraints) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}