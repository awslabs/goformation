package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElasticLoadBalancingV2Listener AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::Listener)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html
type AWSElasticLoadBalancingV2Listener struct {

	// Certificates AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-certificates
	Certificates []AWSElasticLoadBalancingV2Listener_Certificate `json:"Certificates"`

	// DefaultActions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-defaultactions
	DefaultActions []AWSElasticLoadBalancingV2Listener_Action `json:"DefaultActions"`

	// LoadBalancerArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-loadbalancerarn
	LoadBalancerArn string `json:"LoadBalancerArn"`

	// Port AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-port
	Port int64 `json:"Port"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-protocol
	Protocol string `json:"Protocol"`

	// SslPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-sslpolicy
	SslPolicy string `json:"SslPolicy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2Listener) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::Listener"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2Listener) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2ListenerResources retrieves all AWSElasticLoadBalancingV2Listener items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2ListenerResources(template *Template) map[string]*AWSElasticLoadBalancingV2Listener {

	results := map[string]*AWSElasticLoadBalancingV2Listener{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2Listener{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2ListenerWithName retrieves all AWSElasticLoadBalancingV2Listener items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElasticLoadBalancingV2ListenerWithName(name string, template *Template) (*AWSElasticLoadBalancingV2Listener, error) {

	result := &AWSElasticLoadBalancingV2Listener{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2Listener{}, errors.New("resource not found")

}
