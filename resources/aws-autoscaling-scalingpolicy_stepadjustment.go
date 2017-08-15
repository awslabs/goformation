package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::AutoScaling::ScalingPolicy.StepAdjustment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html
type AWSAutoScalingScalingPolicy_StepAdjustment struct {

	// MetricIntervalLowerBound AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-metricintervallowerbound
	MetricIntervalLowerBound float64 `json:"MetricIntervalLowerBound"`

	// MetricIntervalUpperBound AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-metricintervalupperbound
	MetricIntervalUpperBound float64 `json:"MetricIntervalUpperBound"`

	// ScalingAdjustment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-scalingadjustment
	ScalingAdjustment int64 `json:"ScalingAdjustment"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScalingPolicy_StepAdjustment) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScalingPolicy.StepAdjustment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingScalingPolicy_StepAdjustment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingScalingPolicy_StepAdjustmentResources retrieves all AWSAutoScalingScalingPolicy_StepAdjustment items from a CloudFormation template
func GetAllAWSAutoScalingScalingPolicy_StepAdjustment(template *Template) map[string]*AWSAutoScalingScalingPolicy_StepAdjustment {

	results := map[string]*AWSAutoScalingScalingPolicy_StepAdjustment{}
	for name, resource := range template.Resources {
		result := &AWSAutoScalingScalingPolicy_StepAdjustment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingScalingPolicy_StepAdjustmentWithName retrieves all AWSAutoScalingScalingPolicy_StepAdjustment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSAutoScalingScalingPolicy_StepAdjustment(name string, template *Template) (*AWSAutoScalingScalingPolicy_StepAdjustment, error) {

	result := &AWSAutoScalingScalingPolicy_StepAdjustment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingScalingPolicy_StepAdjustment{}, errors.New("resource not found")

}
