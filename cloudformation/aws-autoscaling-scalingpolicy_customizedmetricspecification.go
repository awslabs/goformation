package cloudformation

import (
	"encoding/json"
)

// AWSAutoScalingScalingPolicy_CustomizedMetricSpecification AWS CloudFormation Resource (AWS::AutoScaling::ScalingPolicy.CustomizedMetricSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html
type AWSAutoScalingScalingPolicy_CustomizedMetricSpecification struct {

	// Dimensions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-dimensions
	Dimensions []AWSAutoScalingScalingPolicy_MetricDimension `json:"Dimensions,omitempty"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-metricname
	MetricName *Value `json:"MetricName,omitempty"`

	// Namespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-namespace
	Namespace *Value `json:"Namespace,omitempty"`

	// Statistic AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-statistic
	Statistic *Value `json:"Statistic,omitempty"`

	// Unit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-unit
	Unit *Value `json:"Unit,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScalingPolicy_CustomizedMetricSpecification) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScalingPolicy.CustomizedMetricSpecification"
}

func (r *AWSAutoScalingScalingPolicy_CustomizedMetricSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
