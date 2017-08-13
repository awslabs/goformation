package resources

// AWS::EMR::InstanceFleetConfig.EbsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html
type AWSEMRInstanceFleetConfigEbsConfiguration struct {
    
    // EbsBlockDeviceConfigs AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html#cfn-elasticmapreduce-instancefleetconfig-ebsconfiguration-ebsblockdeviceconfigs
    EbsBlockDeviceConfigs []AWSEMRInstanceFleetConfigEbsConfigurationEbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs"`
    
    // EbsOptimized AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html#cfn-elasticmapreduce-instancefleetconfig-ebsconfiguration-ebsoptimized
    EbsOptimized bool `json:"EbsOptimized"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfigEbsConfiguration) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceFleetConfig.EbsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfigEbsConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}