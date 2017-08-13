package resources

// AWS::EMR::Cluster.SpotProvisioningSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html
type AWSEMRClusterSpotProvisioningSpecification struct {
    
    // BlockDurationMinutes AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html#cfn-elasticmapreduce-cluster-spotprovisioningspecification-blockdurationminutes
    BlockDurationMinutes int64 `json:"BlockDurationMinutes"`
    
    // TimeoutAction AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html#cfn-elasticmapreduce-cluster-spotprovisioningspecification-timeoutaction
    TimeoutAction string `json:"TimeoutAction"`
    
    // TimeoutDurationMinutes AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-spotprovisioningspecification.html#cfn-elasticmapreduce-cluster-spotprovisioningspecification-timeoutdurationminutes
    TimeoutDurationMinutes int64 `json:"TimeoutDurationMinutes"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterSpotProvisioningSpecification) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.SpotProvisioningSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterSpotProvisioningSpecification) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}