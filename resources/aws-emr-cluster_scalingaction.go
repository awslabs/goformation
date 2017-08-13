package resources

// AWS::EMR::Cluster.ScalingAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html
type AWSEMRClusterScalingAction struct {
    
    // Market AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-market
    Market string `json:"Market"`
    
    // SimpleScalingPolicyConfiguration AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-simplescalingpolicyconfiguration
    SimpleScalingPolicyConfiguration AWSEMRClusterScalingActionSimpleScalingPolicyConfiguration `json:"SimpleScalingPolicyConfiguration"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterScalingAction) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.ScalingAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterScalingAction) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}