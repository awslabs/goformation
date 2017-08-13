package resources

// AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html
type AWSEMRInstanceGroupConfigSimpleScalingPolicyConfiguration struct {
    
    // AdjustmentType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-adjustmenttype
    AdjustmentType string `json:"AdjustmentType"`
    
    // CoolDown AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-cooldown
    CoolDown int64 `json:"CoolDown"`
    
    // ScalingAdjustment AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-scalingadjustment
    ScalingAdjustment int64 `json:"ScalingAdjustment"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfigSimpleScalingPolicyConfiguration) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfigSimpleScalingPolicyConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}