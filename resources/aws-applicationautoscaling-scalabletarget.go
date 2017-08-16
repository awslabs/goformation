package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApplicationAutoScalingScalableTarget AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalableTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html
type AWSApplicationAutoScalingScalableTarget struct {

	// MaxCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-maxcapacity
	MaxCapacity int `json:"MaxCapacity"`

	// MinCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-mincapacity
	MinCapacity int `json:"MinCapacity"`

	// ResourceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-resourceid
	ResourceId string `json:"ResourceId"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-rolearn
	RoleARN string `json:"RoleARN"`

	// ScalableDimension AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-scalabledimension
	ScalableDimension string `json:"ScalableDimension"`

	// ServiceNamespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-servicenamespace
	ServiceNamespace string `json:"ServiceNamespace"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalableTarget) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalableTarget"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalableTarget) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApplicationAutoScalingScalableTargetResources retrieves all AWSApplicationAutoScalingScalableTarget items from a CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalableTargetResources() map[string]*AWSApplicationAutoScalingScalableTarget {

	results := map[string]*AWSApplicationAutoScalingScalableTarget{}
	for name, resource := range t.Resources {
		result := &AWSApplicationAutoScalingScalableTarget{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalableTargetWithName retrieves all AWSApplicationAutoScalingScalableTarget items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalableTargetWithName(name string) (*AWSApplicationAutoScalingScalableTarget, error) {

	result := &AWSApplicationAutoScalingScalableTarget{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalableTarget{}, errors.New("resource not found")

}
