package resources

// AWS::EMR::InstanceFleetConfig.InstanceFleetProvisioningSpecifications AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html
type AWSEMRInstanceFleetConfigInstanceFleetProvisioningSpecifications struct {
    
    // SpotSpecification AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications-spotspecification
    SpotSpecification AWSEMRInstanceFleetConfigInstanceFleetProvisioningSpecificationsSpotProvisioningSpecification `json:"SpotSpecification"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfigInstanceFleetProvisioningSpecifications) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceFleetConfig.InstanceFleetProvisioningSpecifications"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfigInstanceFleetProvisioningSpecifications) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}