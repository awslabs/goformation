package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

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

// GetAllAWSApplicationAutoScalingScalingPolicy_MetricDimensionResources retrieves all AWSApplicationAutoScalingScalingPolicy_MetricDimension items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicy_MetricDimension(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy_MetricDimension {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy_MetricDimension{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy_MetricDimension{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicy_MetricDimensionWithName retrieves all AWSApplicationAutoScalingScalingPolicy_MetricDimension items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApplicationAutoScalingScalingPolicy_MetricDimension(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy_MetricDimension, error) {

	result := &AWSApplicationAutoScalingScalingPolicy_MetricDimension{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy_MetricDimension{}, errors.New("resource not found")

}
