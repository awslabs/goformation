package resources

// AWS::EMR::Cluster.InstanceFleetConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html
type AWSEMRClusterInstanceFleetConfig struct {
    
    // InstanceTypeConfigs AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-instancetypeconfigs
    InstanceTypeConfigs []AWSEMRClusterInstanceFleetConfigInstanceTypeConfig `json:"InstanceTypeConfigs"`
    
    // LaunchSpecifications AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-launchspecifications
    LaunchSpecifications AWSEMRClusterInstanceFleetConfigInstanceFleetProvisioningSpecifications `json:"LaunchSpecifications"`
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-name
    Name string `json:"Name"`
    
    // TargetOnDemandCapacity AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetondemandcapacity
    TargetOnDemandCapacity int64 `json:"TargetOnDemandCapacity"`
    
    // TargetSpotCapacity AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetspotcapacity
    TargetSpotCapacity int64 `json:"TargetSpotCapacity"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterInstanceFleetConfig) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.InstanceFleetConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterInstanceFleetConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}