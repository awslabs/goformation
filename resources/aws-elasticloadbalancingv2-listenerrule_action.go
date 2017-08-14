package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::ListenerRule.Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type AWSElasticLoadBalancingV2ListenerRule_Action struct {

	// TargetGroupArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-targetgrouparn

	TargetGroupArn string `json:"TargetGroupArn"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html#cfn-elasticloadbalancingv2-listener-actions-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule.Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2ListenerRule_Action) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2ListenerRule_ActionResources retrieves all AWSElasticLoadBalancingV2ListenerRule_Action items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2ListenerRule_Action(template *Template) map[string]*AWSElasticLoadBalancingV2ListenerRule_Action {

	results := map[string]*AWSElasticLoadBalancingV2ListenerRule_Action{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2ListenerRule_Action{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2ListenerRule_ActionWithName retrieves all AWSElasticLoadBalancingV2ListenerRule_Action items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2ListenerRule_Action(name string, template *Template) (*AWSElasticLoadBalancingV2ListenerRule_Action, error) {

	result := &AWSElasticLoadBalancingV2ListenerRule_Action{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2ListenerRule_Action{}, errors.New("resource not found")

}
