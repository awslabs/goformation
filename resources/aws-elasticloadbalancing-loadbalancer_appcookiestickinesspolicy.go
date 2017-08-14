package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html
type AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy struct {

	// CookieName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html#cfn-elb-appcookiestickinesspolicy-cookiename

	CookieName string `json:"CookieName"`

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html#cfn-elb-appcookiestickinesspolicy-policyname

	PolicyName string `json:"PolicyName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicyResources retrieves all AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicyWithName retrieves all AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy{}, errors.New("resource not found")

}
