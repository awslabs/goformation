package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html
type AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification struct {

	// PredefinedMetricType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predefinedmetricspecification-predefinedmetrictype
	PredefinedMetricType string `json:"PredefinedMetricType"`

	// ResourceLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predefinedmetricspecification-resourcelabel
	ResourceLabel string `json:"ResourceLabel"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecificationResources retrieves all AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecificationWithName retrieves all AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification, error) {

	result := &AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy_PredefinedMetricSpecification{}, errors.New("resource not found")

}
