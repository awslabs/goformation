package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::WebACL.Rule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html
type AWSWAFRegionalWebACL_Rule struct {

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-action

	Action AWSWAFRegionalWebACL_Action `json:"Action"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-priority

	Priority int64 `json:"Priority"`

	// RuleId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-ruleid

	RuleId string `json:"RuleId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACL_Rule) AWSCloudFormationType() string {
	return "AWS::WAFRegional::WebACL.Rule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalWebACL_Rule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalWebACL_RuleResources retrieves all AWSWAFRegionalWebACL_Rule items from a CloudFormation template
func GetAllAWSWAFRegionalWebACL_Rule(template *Template) map[string]*AWSWAFRegionalWebACL_Rule {

	results := map[string]*AWSWAFRegionalWebACL_Rule{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalWebACL_Rule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalWebACL_RuleWithName retrieves all AWSWAFRegionalWebACL_Rule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalWebACL_Rule(name string, template *Template) (*AWSWAFRegionalWebACL_Rule, error) {

	result := &AWSWAFRegionalWebACL_Rule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalWebACL_Rule{}, errors.New("resource not found")

}
