package resources

// AWS::OpsWorks::Layer.AutoScalingThresholds AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html
type AWSOpsWorksLayerAutoScalingThresholds struct {
    
    // CpuThreshold AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-cputhreshold
    CpuThreshold float64 `json:"CpuThreshold"`
    
    // IgnoreMetricsTime AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-ignoremetricstime
    IgnoreMetricsTime int64 `json:"IgnoreMetricsTime"`
    
    // InstanceCount AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-instancecount
    InstanceCount int64 `json:"InstanceCount"`
    
    // LoadThreshold AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-loadthreshold
    LoadThreshold float64 `json:"LoadThreshold"`
    
    // MemoryThreshold AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-memorythreshold
    MemoryThreshold float64 `json:"MemoryThreshold"`
    
    // ThresholdsWaitTime AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-thresholdwaittime
    ThresholdsWaitTime int64 `json:"ThresholdsWaitTime"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayerAutoScalingThresholds) AWSCloudFormationType() string {
    return "AWS::OpsWorks::Layer.AutoScalingThresholds"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayerAutoScalingThresholds) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}