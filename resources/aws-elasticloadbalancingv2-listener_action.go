package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::Listener.Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html
type AWSElasticLoadBalancingV2Listener_Action struct {

	// TargetGroupArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html#cfn-elasticloadbalancingv2-listener-defaultactions-targetgrouparn
	TargetGroupArn string `json:"TargetGroupArn"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html#cfn-elasticloadbalancingv2-listener-defaultactions-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2Listener_Action) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::Listener.Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2Listener_Action) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2Listener_ActionResources retrieves all AWSElasticLoadBalancingV2Listener_Action items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2Listener_Action(template *Template) map[string]*AWSElasticLoadBalancingV2Listener_Action {

	results := map[string]*AWSElasticLoadBalancingV2Listener_Action{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2Listener_Action{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2Listener_ActionWithName retrieves all AWSElasticLoadBalancingV2Listener_Action items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2Listener_Action(name string, template *Template) (*AWSElasticLoadBalancingV2Listener_Action, error) {

	result := &AWSElasticLoadBalancingV2Listener_Action{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2Listener_Action{}, errors.New("resource not found")

}
