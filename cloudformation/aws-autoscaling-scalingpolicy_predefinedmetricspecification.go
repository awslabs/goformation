package cloudformation

// AWSAutoScalingScalingPolicy_PredefinedMetricSpecification AWS CloudFormation Resource (AWS::AutoScaling::ScalingPolicy.PredefinedMetricSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html
type AWSAutoScalingScalingPolicy_PredefinedMetricSpecification struct {

	// PredefinedMetricType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-autoscaling-scalingpolicy-predefinedmetricspecification-predefinedmetrictype
	PredefinedMetricType string `json:"PredefinedMetricType,omitempty"`

	// ResourceLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-autoscaling-scalingpolicy-predefinedmetricspecification-resourcelabel
	ResourceLabel string `json:"ResourceLabel,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScalingPolicy_PredefinedMetricSpecification) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScalingPolicy.PredefinedMetricSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingScalingPolicy_PredefinedMetricSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.5.0"
}
