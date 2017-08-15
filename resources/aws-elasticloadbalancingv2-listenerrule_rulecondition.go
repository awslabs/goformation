package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticLoadBalancingV2::ListenerRule.RuleCondition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html
type AWSElasticLoadBalancingV2ListenerRule_RuleCondition struct {

	// Field AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-conditions-field

	Field string `json:"Field"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-conditions-values

	Values []string `json:"Values"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2ListenerRule_RuleCondition) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::ListenerRule.RuleCondition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2ListenerRule_RuleCondition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticLoadBalancingV2ListenerRule_RuleConditionResources retrieves all AWSElasticLoadBalancingV2ListenerRule_RuleCondition items from a CloudFormation template
func GetAllAWSElasticLoadBalancingV2ListenerRule_RuleCondition(template *Template) map[string]*AWSElasticLoadBalancingV2ListenerRule_RuleCondition {

	results := map[string]*AWSElasticLoadBalancingV2ListenerRule_RuleCondition{}
	for name, resource := range template.Resources {
		result := &AWSElasticLoadBalancingV2ListenerRule_RuleCondition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticLoadBalancingV2ListenerRule_RuleConditionWithName retrieves all AWSElasticLoadBalancingV2ListenerRule_RuleCondition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticLoadBalancingV2ListenerRule_RuleCondition(name string, template *Template) (*AWSElasticLoadBalancingV2ListenerRule_RuleCondition, error) {

	result := &AWSElasticLoadBalancingV2ListenerRule_RuleCondition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticLoadBalancingV2ListenerRule_RuleCondition{}, errors.New("resource not found")

}
