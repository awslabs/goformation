package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html
type AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy struct {

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html#cfn-elb-connectiondrainingpolicy-enabled

	Enabled bool `json:"Enabled"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html#cfn-elb-connectiondrainingpolicy-timeout

	Timeout int64 `json:"Timeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicyResources retrieves all AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicyWithName retrieves all AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy{}, errors.New("resource not found")

}
