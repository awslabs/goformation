package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApplicationAutoScaling::ScalingPolicy.CustomizedMetricSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html
type AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification struct {

	// Dimensions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-customizedmetricspecification-dimensions
	Dimensions []AWSApplicationAutoScalingScalingPolicy_MetricDimension `json:"Dimensions"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-customizedmetricspecification-metricname
	MetricName string `json:"MetricName"`

	// Namespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-customizedmetricspecification-namespace
	Namespace string `json:"Namespace"`

	// Statistic AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-customizedmetricspecification-statistic
	Statistic string `json:"Statistic"`

	// Unit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-customizedmetricspecification-unit
	Unit string `json:"Unit"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.CustomizedMetricSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecificationResources retrieves all AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecificationWithName retrieves all AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification, error) {

	result := &AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy_CustomizedMetricSpecification{}, errors.New("resource not found")

}
