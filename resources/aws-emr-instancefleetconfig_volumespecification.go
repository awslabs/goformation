package resources

// AWS::EMR::InstanceFleetConfig.VolumeSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html
type AWSEMRInstanceFleetConfigVolumeSpecification struct {
    
    // Iops AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-iops
    Iops int64 `json:"Iops"`
    
    // SizeInGB AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-sizeingb
    SizeInGB int64 `json:"SizeInGB"`
    
    // VolumeType AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-volumetype
    VolumeType string `json:"VolumeType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfigVolumeSpecification) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceFleetConfig.VolumeSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfigVolumeSpecification) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}