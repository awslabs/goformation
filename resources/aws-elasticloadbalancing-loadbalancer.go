package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
type AWSElasticLoadBalancingLoadBalancer struct {

	// AccessLoggingPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-accessloggingpolicy
	AccessLoggingPolicy AWSElasticLoadBalancingLoadBalancer_AccessLoggingPolicy `json:"AccessLoggingPolicy"`

	// AppCookieStickinessPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-appcookiestickinesspolicy
	AppCookieStickinessPolicy []AWSElasticLoadBalancingLoadBalancer_AppCookieStickinessPolicy `json:"AppCookieStickinessPolicy"`

	// AvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-availabilityzones
	AvailabilityZones []string `json:"AvailabilityZones"`

	// ConnectionDrainingPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-connectiondrainingpolicy
	ConnectionDrainingPolicy AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy `json:"ConnectionDrainingPolicy"`

	// ConnectionSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-connectionsettings
	ConnectionSettings AWSElasticLoadBalancingLoadBalancer_ConnectionSettings `json:"ConnectionSettings"`

	// CrossZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-crosszone
	CrossZone bool `json:"CrossZone"`

	// HealthCheck AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-healthcheck
	HealthCheck AWSElasticLoadBalancingLoadBalancer_HealthCheck `json:"HealthCheck"`

	// Instances AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-instances
	Instances []string `json:"Instances"`

	// LBCookieStickinessPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-lbcookiestickinesspolicy
	LBCookieStickinessPolicy []AWSElasticLoadBalancingLoadBalancer_LBCookieStickinessPolicy `json:"LBCookieStickinessPolicy"`

	// Listeners AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-listeners
	Listeners []AWSElasticLoadBalancingLoadBalancer_Listeners `json:"Listeners"`

	// LoadBalancerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-elbname
	LoadBalancerName string `json:"LoadBalancerName"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-policies
	Policies []AWSElasticLoadBalancingLoadBalancer_Policies `json:"Policies"`

	// Scheme AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-scheme
	Scheme string `json:"Scheme"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-securitygroups
	SecurityGroups []string `json:"SecurityGroups"`

	// Subnets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-subnets
	Subnets []string `json:"Subnets"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-elasticloadbalancing-loadbalancer-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancerResources retrieves all AWSElasticLoadBalancingLoadBalancer items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancerResources(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancerWithName retrieves all AWSElasticLoadBalancingLoadBalancer items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElasticLoadBalancingLoadBalancerWithName(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer, error) {

	result := &AWSElasticLoadBalancingLoadBalancer{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer{}, errors.New("resource not found")

}
