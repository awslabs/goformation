package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html
type AWSElasticLoadBalancingLoadBalancer_ConnectionSettings struct {

	// IdleTimeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html#cfn-elb-connectionsettings-idletimeout

	IdleTimeout int64 `json:"IdleTimeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionSettings) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingLoadBalancer_ConnectionSettingsResources retrieves all AWSElasticLoadBalancingLoadBalancer_ConnectionSettings items from a CloudFormation template
func GetAllAWSElasticLoadBalancingLoadBalancer_ConnectionSettings(template *Template) map[string]*AWSElasticLoadBalancingLoadBalancer_ConnectionSettings {

	results := map[string]*AWSElasticLoadBalancingLoadBalancer_ConnectionSettings{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingLoadBalancer_ConnectionSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingLoadBalancer_ConnectionSettingsWithName retrieves all AWSElasticLoadBalancingLoadBalancer_ConnectionSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingLoadBalancer_ConnectionSettings(name string, template *Template) (*AWSElasticLoadBalancingLoadBalancer_ConnectionSettings, error) {

	result := &AWSElasticLoadBalancingLoadBalancer_ConnectionSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingLoadBalancer_ConnectionSettings{}, errors.New("resource not found")

}
