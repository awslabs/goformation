package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.Listeners AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
type AWSElasticLoadBalancingLoadBalancer_Listeners struct {

	// InstancePort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceport
	InstancePort string `json:"InstancePort"`

	// InstanceProtocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceprotocol
	InstanceProtocol string `json:"InstanceProtocol"`

	// LoadBalancerPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-loadbalancerport
	LoadBalancerPort string `json:"LoadBalancerPort"`

	// PolicyNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames
	PolicyNames []string `json:"PolicyNames"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-protocol
	Protocol string `json:"Protocol"`

	// SSLCertificateId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-sslcertificateid
	SSLCertificateId string `json:"SSLCertificateId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_Listeners) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.Listeners"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_Listeners) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_ListenersResources retrieves all AWSElasticLoadBalancingLoadBalancer_Listeners items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_Listeners(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_Listeners {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_Listeners{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_Listeners{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_ListenersWithName retrieves all AWSElasticLoadBalancingLoadBalancer_Listeners items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_Listeners(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_Listeners, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_Listeners{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_Listeners{}, errors.New("resource not found")

}
