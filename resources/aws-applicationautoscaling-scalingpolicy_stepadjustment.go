package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

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

// GetAllAWSApplicationAutoScalingScalingPolicy_StepAdjustmentResources retrieves all AWSApplicationAutoScalingScalingPolicy_StepAdjustment items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicy_StepAdjustment(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy_StepAdjustment {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy_StepAdjustment{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy_StepAdjustment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicy_StepAdjustmentWithName retrieves all AWSApplicationAutoScalingScalingPolicy_StepAdjustment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApplicationAutoScalingScalingPolicy_StepAdjustment(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy_StepAdjustment, error) {

	result := &AWSApplicationAutoScalingScalingPolicy_StepAdjustment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy_StepAdjustment{}, errors.New("resource not found")

}
