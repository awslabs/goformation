package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingScalingPolicyConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html
type AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration struct {

	// CustomizedMetricSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-customizedmetricspecification

	CustomizedMetricSpecification AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification `json:"CustomizedMetricSpecification"`

	// PredefinedMetricSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-predefinedmetricspecification

	PredefinedMetricSpecification AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification `json:"PredefinedMetricSpecification"`

	// ScaleInCooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-scaleincooldown

	ScaleInCooldown int64 `json:"ScaleInCooldown"`

	// ScaleOutCooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-scaleoutcooldown

	ScaleOutCooldown int64 `json:"ScaleOutCooldown"`

	// TargetValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-targetvalue

	TargetValue float64 `json:"TargetValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingScalingPolicyConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfigurationResources retrieves all AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfigurationWithName retrieves all AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration, error) {

	result := &AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration{}, errors.New("resource not found")

}
