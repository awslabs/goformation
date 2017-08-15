package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::AutoScaling::AutoScalingGroup.TagProperty AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html
type AWSAutoScalingAutoScalingGroup_TagProperty struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-Key

	Key string `json:"Key"`

	// PropagateAtLaunch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-PropagateAtLaunch

	PropagateAtLaunch bool `json:"PropagateAtLaunch"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-Value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_TagProperty) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.TagProperty"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingAutoScalingGroup_TagProperty) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSAutoScalingAutoScalingGroup_TagPropertyResources retrieves all AWSAutoScalingAutoScalingGroup_TagProperty items from a CloudFormation template
func GetAllAWSAutoScalingAutoScalingGroup_TagProperty(template *Template) map[string]*AWSAutoScalingAutoScalingGroup_TagProperty {

	results := map[string]*AWSAutoScalingAutoScalingGroup_TagProperty{}
	for name, resource := range template.Resources {
		result := &AWSAutoScalingAutoScalingGroup_TagProperty{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSAutoScalingAutoScalingGroup_TagPropertyWithName retrieves all AWSAutoScalingAutoScalingGroup_TagProperty items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSAutoScalingAutoScalingGroup_TagProperty(name string, template *Template) (*AWSAutoScalingAutoScalingGroup_TagProperty, error) {

	result := &AWSAutoScalingAutoScalingGroup_TagProperty{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSAutoScalingAutoScalingGroup_TagProperty{}, errors.New("resource not found")

}
