package resources

// AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html
type AWSEMRClusterInstanceFleetProvisioningSpecifications struct {
    
    // SpotSpecification AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-cluster-instancefleetprovisioningspecifications-spotspecification
    SpotSpecification AWSEMRClusterInstanceFleetProvisioningSpecificationsSpotProvisioningSpecification `json:"SpotSpecification"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterInstanceFleetProvisioningSpecifications) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.InstanceFleetProvisioningSpecifications"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterInstanceFleetProvisioningSpecifications) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}