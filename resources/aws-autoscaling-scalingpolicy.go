package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSAutoScalingScalingPolicy AWS CloudFormation Resource (AWS::AutoScaling::ScalingPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html
type AWSAutoScalingScalingPolicy struct {

	// AdjustmentType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-adjustmenttype
	AdjustmentType string `json:"AdjustmentType"`

	// AutoScalingGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-autoscalinggroupname
	AutoScalingGroupName string `json:"AutoScalingGroupName"`

	// Cooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-cooldown
	Cooldown string `json:"Cooldown"`

	// EstimatedInstanceWarmup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-estimatedinstancewarmup
	EstimatedInstanceWarmup int `json:"EstimatedInstanceWarmup"`

	// MetricAggregationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-metricaggregationtype
	MetricAggregationType string `json:"MetricAggregationType"`

	// MinAdjustmentMagnitude AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-minadjustmentmagnitude
	MinAdjustmentMagnitude int `json:"MinAdjustmentMagnitude"`

	// PolicyType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-policytype
	PolicyType string `json:"PolicyType"`

	// ScalingAdjustment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-scalingadjustment
	ScalingAdjustment int `json:"ScalingAdjustment"`

	// StepAdjustments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-stepadjustments
	StepAdjustments []AWSAutoScalingScalingPolicy_StepAdjustment `json:"StepAdjustments"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScalingPolicy) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScalingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingScalingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingScalingPolicyResources retrieves all AWSAutoScalingScalingPolicy items from a CloudFormation template
func (t *Template) GetAllAWSAutoScalingScalingPolicyResources() map[string]*AWSAutoScalingScalingPolicy {

	results := map[string]*AWSAutoScalingScalingPolicy{}
	for name, resource := range t.Resources {
		result := &AWSAutoScalingScalingPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingScalingPolicyWithName retrieves all AWSAutoScalingScalingPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScalingPolicyWithName(name string) (*AWSAutoScalingScalingPolicy, error) {

	result := &AWSAutoScalingScalingPolicy{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingScalingPolicy{}, errors.New("resource not found")

}
