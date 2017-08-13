package resources

// AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
type AWSEMRInstanceGroupConfigEbsBlockDeviceConfig struct {
    
    // VolumeSpecification AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification
    VolumeSpecification AWSEMRInstanceGroupConfigEbsBlockDeviceConfigVolumeSpecification `json:"VolumeSpecification"`
    
    // VolumesPerInstance AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance
    VolumesPerInstance int64 `json:"VolumesPerInstance"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfigEbsBlockDeviceConfig) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfigEbsBlockDeviceConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}