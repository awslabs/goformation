package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html
type AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute struct {

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes-key

	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2TargetGroup_TargetGroupAttributeResources retrieves all AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute(template *Template) map[string]*AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute {

	results := map[string]*AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2TargetGroup_TargetGroupAttributeWithName retrieves all AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute(name string, template *Template) (*AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute, error) {

	result := &AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2TargetGroup_TargetGroupAttribute{}, errors.New("resource not found")

}
