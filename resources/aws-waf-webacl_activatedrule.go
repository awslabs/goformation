package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::WebACL.ActivatedRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html
type AWSWAFWebACL_ActivatedRule struct {

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-action
	Action AWSWAFWebACL_WafAction `json:"Action"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-priority
	Priority int64 `json:"Priority"`

	// RuleId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html#cfn-waf-webacl-rules-ruleid
	RuleId string `json:"RuleId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACL_ActivatedRule) AWSCloudFormationType() string {
	return "AWS::WAF::WebACL.ActivatedRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFWebACL_ActivatedRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFWebACL_ActivatedRuleResources retrieves all AWSWAFWebACL_ActivatedRule items from a CloudFormation template
func GetAllAWSWAFWebACL_ActivatedRule(template *Template) map[string]*AWSWAFWebACL_ActivatedRule {

	results := map[string]*AWSWAFWebACL_ActivatedRule{}
	for name, resource := range template.Resources {
		result := &AWSWAFWebACL_ActivatedRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFWebACL_ActivatedRuleWithName retrieves all AWSWAFWebACL_ActivatedRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFWebACL_ActivatedRule(name string, template *Template) (*AWSWAFWebACL_ActivatedRule, error) {

	result := &AWSWAFWebACL_ActivatedRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFWebACL_ActivatedRule{}, errors.New("resource not found")

}
