package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.LoadBasedAutoScaling AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html
type AWSOpsWorksLayer_LoadBasedAutoScaling struct {

	// DownScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-downscaling
	DownScaling AWSOpsWorksLayer_AutoScalingThresholds `json:"DownScaling"`

	// Enable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-enable
	Enable bool `json:"Enable"`

	// UpScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-upscaling
	UpScaling AWSOpsWorksLayer_AutoScalingThresholds `json:"UpScaling"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_LoadBasedAutoScaling) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.LoadBasedAutoScaling"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_LoadBasedAutoScaling) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksLayer_LoadBasedAutoScalingResources retrieves all AWSOpsWorksLayer_LoadBasedAutoScaling items from a CloudFormation template
func GetAllAWSOpsWorksLayer_LoadBasedAutoScaling(template *Template) map[string]*AWSOpsWorksLayer_LoadBasedAutoScaling {

	results := map[string]*AWSOpsWorksLayer_LoadBasedAutoScaling{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksLayer_LoadBasedAutoScaling{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksLayer_LoadBasedAutoScalingWithName retrieves all AWSOpsWorksLayer_LoadBasedAutoScaling items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksLayer_LoadBasedAutoScaling(name string, template *Template) (*AWSOpsWorksLayer_LoadBasedAutoScaling, error) {

	result := &AWSOpsWorksLayer_LoadBasedAutoScaling{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksLayer_LoadBasedAutoScaling{}, errors.New("resource not found")

}
