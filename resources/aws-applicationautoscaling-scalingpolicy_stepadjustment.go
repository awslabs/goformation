package resources

// AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustment.html
type AWSApplicationAutoScalingScalingPolicy_StepAdjustment struct {

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
func (r *AWSApplicationAutoScalingScalingPolicy_StepAdjustment) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.StepAdjustment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy_StepAdjustment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
