package resources

// AWS::EMR::Cluster.ScalingTrigger AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html
type AWSEMRClusterScalingTrigger struct {
    
    // CloudWatchAlarmDefinition AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html#cfn-elasticmapreduce-cluster-scalingtrigger-cloudwatchalarmdefinition
    CloudWatchAlarmDefinition AWSEMRClusterScalingTriggerCloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterScalingTrigger) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.ScalingTrigger"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterScalingTrigger) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}