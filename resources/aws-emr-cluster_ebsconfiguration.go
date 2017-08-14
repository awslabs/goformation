package resources

// AWS::EMR::Cluster.EbsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
type AWSEMRClusterEbsConfiguration struct {

	// EbsBlockDeviceConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs []AWSEMRClusterEbsConfigurationEbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterEbsConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.EbsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterEbsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
