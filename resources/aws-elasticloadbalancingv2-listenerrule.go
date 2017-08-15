package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::ListenerRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
type AWSElasticLoadBalancingV2ListenerRule struct {

	// Actions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-actions
	Actions []AWSElasticLoadBalancingV2ListenerRule_Action `json:"Actions"`

	// Conditions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-conditions
	Conditions []AWSElasticLoadBalancingV2ListenerRule_RuleCondition `json:"Conditions"`

	// ListenerArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-listenerarn
	ListenerArn string `json:"ListenerArn"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html#cfn-elasticloadbalancingv2-listenerrule-priority
	Priority int64 `json:"Priority"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2ListenerRule) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2ListenerRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2ListenerRuleResources retrieves all AWSElasticLoadBalancingV2ListenerRule items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2ListenerRule(template *Template) map[string]*AWSElasticLoadBalancingV2ListenerRule {

	results := map[string]*AWSElasticLoadBalancingV2ListenerRule{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2ListenerRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2ListenerRuleWithName retrieves all AWSElasticLoadBalancingV2ListenerRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2ListenerRule(name string, template *Template) (*AWSElasticLoadBalancingV2ListenerRule, error) {

	result := &AWSElasticLoadBalancingV2ListenerRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2ListenerRule{}, errors.New("resource not found")

}
