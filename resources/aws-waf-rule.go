package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::Rule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html
type AWSWAFRule struct {

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-metricname
	MetricName string `json:"MetricName"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-name
	Name string `json:"Name"`

	// Predicates AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-predicates
	Predicates []AWSWAFRule_Predicate `json:"Predicates"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRule) AWSCloudFormationType() string {
	return "AWS::WAF::Rule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRuleResources retrieves all AWSWAFRule items from a CloudFormation template
func GetAllAWSWAFRule(template *Template) map[string]*AWSWAFRule {

	results := map[string]*AWSWAFRule{}
	for name, resource := range template.Resources {
		result := &AWSWAFRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRuleWithName retrieves all AWSWAFRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRule(name string, template *Template) (*AWSWAFRule, error) {

	result := &AWSWAFRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRule{}, errors.New("resource not found")

}
