package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::Listener.Certificate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-certificates.html
type AWSElasticLoadBalancingV2Listener_Certificate struct {

	// CertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-certificates.html#cfn-elasticloadbalancingv2-listener-certificates-certificatearn

	CertificateArn string `json:"CertificateArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2Listener_Certificate) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::Listener.Certificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2Listener_Certificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2Listener_CertificateResources retrieves all AWSElasticLoadBalancingV2Listener_Certificate items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2Listener_Certificate(template *Template) map[string]*AWSElasticLoadBalancingV2Listener_Certificate {

	results := map[string]*AWSElasticLoadBalancingV2Listener_Certificate{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2Listener_Certificate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2Listener_CertificateWithName retrieves all AWSElasticLoadBalancingV2Listener_Certificate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2Listener_Certificate(name string, template *Template) (*AWSElasticLoadBalancingV2Listener_Certificate, error) {

	result := &AWSElasticLoadBalancingV2Listener_Certificate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2Listener_Certificate{}, errors.New("resource not found")

}
