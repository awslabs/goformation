package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.Policies AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html
type AWSElasticLoadBalancingLoadBalancer_Policies struct {

	// Attributes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-attributes
	Attributes []interface{} `json:"Attributes"`

	// InstancePorts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-instanceports
	InstancePorts []string `json:"InstancePorts"`

	// LoadBalancerPorts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-loadbalancerports
	LoadBalancerPorts []string `json:"LoadBalancerPorts"`

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-policyname
	PolicyName string `json:"PolicyName"`

	// PolicyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html#cfn-ec2-elb-policy-policytype
	PolicyType string `json:"PolicyType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_Policies) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.Policies"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_Policies) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_PoliciesResources retrieves all AWSElasticLoadBalancingLoadBalancer_Policies items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_Policies(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_Policies {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_Policies{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_Policies{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_PoliciesWithName retrieves all AWSElasticLoadBalancingLoadBalancer_Policies items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_Policies(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_Policies, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_Policies{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_Policies{}, errors.New("resource not found")

}
