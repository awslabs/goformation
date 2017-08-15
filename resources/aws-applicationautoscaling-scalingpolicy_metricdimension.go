package resources

// AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html
type AWSApplicationAutoScalingScalingPolicy_MetricDimension struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html#cfn-applicationautoscaling-scalingpolicy-metricdimension-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html#cfn-applicationautoscaling-scalingpolicy-metricdimension-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
