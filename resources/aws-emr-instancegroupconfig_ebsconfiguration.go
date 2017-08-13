package resources

// AWS::EMR::InstanceGroupConfig.EbsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
type AWSEMRInstanceGroupConfigEbsConfiguration struct {
    
    // EbsBlockDeviceConfigs AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs
    EbsBlockDeviceConfigs []AWSEMRInstanceGroupConfigEbsConfigurationEbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs"`
    
    // EbsOptimized AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized
    EbsOptimized bool `json:"EbsOptimized"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfigEbsConfiguration) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceGroupConfig.EbsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfigEbsConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}