package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.AutoScalingThresholds AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html
type AWSOpsWorksLayer_AutoScalingThresholds struct {

	// CpuThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-cputhreshold
	CpuThreshold float64 `json:"CpuThreshold"`

	// IgnoreMetricsTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-ignoremetricstime
	IgnoreMetricsTime int64 `json:"IgnoreMetricsTime"`

	// InstanceCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-instancecount
	InstanceCount int64 `json:"InstanceCount"`

	// LoadThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-loadthreshold
	LoadThreshold float64 `json:"LoadThreshold"`

	// MemoryThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-memorythreshold
	MemoryThreshold float64 `json:"MemoryThreshold"`

	// ThresholdsWaitTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling-autoscalingthresholds.html#cfn-opsworks-layer-loadbasedautoscaling-autoscalingthresholds-thresholdwaittime
	ThresholdsWaitTime int64 `json:"ThresholdsWaitTime"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_AutoScalingThresholds) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.AutoScalingThresholds"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_AutoScalingThresholds) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksLayer_AutoScalingThresholdsResources retrieves all AWSOpsWorksLayer_AutoScalingThresholds items from a CloudFormation template
func GetAllAWSOpsWorksLayer_AutoScalingThresholds(template *Template) map[string]*AWSOpsWorksLayer_AutoScalingThresholds {

	results := map[string]*AWSOpsWorksLayer_AutoScalingThresholds{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksLayer_AutoScalingThresholds{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksLayer_AutoScalingThresholdsWithName retrieves all AWSOpsWorksLayer_AutoScalingThresholds items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksLayer_AutoScalingThresholds(name string, template *Template) (*AWSOpsWorksLayer_AutoScalingThresholds, error) {

	result := &AWSOpsWorksLayer_AutoScalingThresholds{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksLayer_AutoScalingThresholds{}, errors.New("resource not found")

}
