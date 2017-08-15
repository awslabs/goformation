package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.LBCookieStickinessPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html
type AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy struct {

	// CookieExpirationPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html#cfn-elb-lbcookiestickinesspolicy-cookieexpirationperiod
	CookieExpirationPeriod string `json:"CookieExpirationPeriod"`

	// PolicyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html#cfn-elb-lbcookiestickinesspolicy-policyname
	PolicyName string `json:"PolicyName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.LBCookieStickinessPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicyResources retrieves all AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicyWithName retrieves all AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy{}, errors.New("resource not found")

}
