package resources

// AWS::EMR::Cluster.EbsBlockDeviceConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
type AWSEMRClusterEbsBlockDeviceConfig struct {
    
    // VolumeSpecification AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification
    VolumeSpecification AWSEMRClusterEbsBlockDeviceConfigVolumeSpecification `json:"VolumeSpecification"`
    
    // VolumesPerInstance AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance
    VolumesPerInstance int64 `json:"VolumesPerInstance"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterEbsBlockDeviceConfig) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.EbsBlockDeviceConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterEbsBlockDeviceConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}