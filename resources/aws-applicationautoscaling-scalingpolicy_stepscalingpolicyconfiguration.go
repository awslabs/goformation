package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApplicationAutoScaling::ScalingPolicy.StepScalingPolicyConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html
type AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration struct {

	// AdjustmentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-adjustmenttype
	AdjustmentType string `json:"AdjustmentType"`

	// Cooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-cooldown
	Cooldown int64 `json:"Cooldown"`

	// MetricAggregationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-metricaggregationtype
	MetricAggregationType string `json:"MetricAggregationType"`

	// MinAdjustmentMagnitude AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-minadjustmentmagnitude
	MinAdjustmentMagnitude int64 `json:"MinAdjustmentMagnitude"`

	// StepAdjustments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustments
	StepAdjustments []AWSApplicationAutoScalingScalingPolicy_StepAdjustment `json:"StepAdjustments"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.StepScalingPolicyConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfigurationResources retrieves all AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfigurationWithName retrieves all AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration, error) {

	result := &AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration{}, errors.New("resource not found")

}
