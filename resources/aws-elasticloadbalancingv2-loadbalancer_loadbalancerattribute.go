package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html
type AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute struct {

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes-key

	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttributeResources retrieves all AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute(template *Template) map[string]*AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute {

	results := map[string]*AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttributeWithName retrieves all AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute(name string, template *Template) (*AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute, error) {

	result := &AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute{}, errors.New("resource not found")

}
