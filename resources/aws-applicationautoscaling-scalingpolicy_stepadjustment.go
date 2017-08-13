package resources

// AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html
type AWSApplicationAutoScalingScalingPolicyStepAdjustment struct {
    
    // MetricIntervalLowerBound AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment-metricintervallowerbound
    MetricIntervalLowerBound float64 `json:"MetricIntervalLowerBound"`
    
    // MetricIntervalUpperBound AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment-metricintervalupperbound
    MetricIntervalUpperBound float64 `json:"MetricIntervalUpperBound"`
    
    // ScalingAdjustment AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment-scalingadjustment
    ScalingAdjustment int64 `json:"ScalingAdjustment"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicyStepAdjustment) AWSCloudFormationType() string {
    return "AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicyStepAdjustment) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}