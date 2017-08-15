package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::LoadBalancer AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html
type AWSElasticLoadBalancingV2LoadBalancer struct {

	// IpAddressType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-ipaddresstype
	IpAddressType string `json:"IpAddressType"`

	// LoadBalancerAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes
	LoadBalancerAttributes []AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute `json:"LoadBalancerAttributes"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-name
	Name string `json:"Name"`

	// Scheme AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-scheme
	Scheme string `json:"Scheme"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-securitygroups
	SecurityGroups []string `json:"SecurityGroups"`

	// Subnets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-subnets
	Subnets []string `json:"Subnets"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2LoadBalancer) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::LoadBalancer"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2LoadBalancer) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2LoadBalancerResources retrieves all AWSElasticLoadBalancingV2LoadBalancer items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2LoadBalancer(template *Template) map[string]*AWSElasticLoadBalancingV2LoadBalancer {

	results := map[string]*AWSElasticLoadBalancingV2LoadBalancer{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2LoadBalancer{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2LoadBalancerWithName retrieves all AWSElasticLoadBalancingV2LoadBalancer items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2LoadBalancer(name string, template *Template) (*AWSElasticLoadBalancingV2LoadBalancer, error) {

	result := &AWSElasticLoadBalancingV2LoadBalancer{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2LoadBalancer{}, errors.New("resource not found")

}
