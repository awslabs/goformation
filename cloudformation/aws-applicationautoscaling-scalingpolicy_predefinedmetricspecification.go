package cloudformation

import (
	"encoding/json"
)

// AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html
type AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification struct {

	// PredefinedMetricType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predefinedmetricspecification-predefinedmetrictype
	PredefinedMetricType *Value `json:"PredefinedMetricType,omitempty"`

	// ResourceLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predefinedmetricspecification-resourcelabel
	ResourceLabel *Value `json:"ResourceLabel,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification"
}

func (r *AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
