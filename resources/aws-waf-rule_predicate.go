package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::Rule.Predicate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html
type AWSWAFRule_Predicate struct {

	// DataId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-dataid
	DataId string `json:"DataId"`

	// Negated AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-negated
	Negated bool `json:"Negated"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html#cfn-waf-rule-predicates-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRule_Predicate) AWSCloudFormationType() string {
	return "AWS::WAF::Rule.Predicate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRule_Predicate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRule_PredicateResources retrieves all AWSWAFRule_Predicate items from a CloudFormation template
func GetAllAWSWAFRule_Predicate(template *Template) map[string]*AWSWAFRule_Predicate {

	results := map[string]*AWSWAFRule_Predicate{}
	for name, resource := range template.Resources {
		result := &AWSWAFRule_Predicate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRule_PredicateWithName retrieves all AWSWAFRule_Predicate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRule_Predicate(name string, template *Template) (*AWSWAFRule_Predicate, error) {

	result := &AWSWAFRule_Predicate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRule_Predicate{}, errors.New("resource not found")

}
