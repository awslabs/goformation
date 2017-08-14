package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::Rule.Predicate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html
type AWSWAFRegionalRule_Predicate struct {

	// DataId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-dataid

	DataId string `json:"DataId"`

	// Negated AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-negated

	Negated bool `json:"Negated"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-rule-predicate.html#cfn-wafregional-rule-predicate-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalRule_Predicate) AWSCloudFormationType() string {
	return "AWS::WAFRegional::Rule.Predicate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalRule_Predicate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalRule_PredicateResources retrieves all AWSWAFRegionalRule_Predicate items from a CloudFormation template
func GetAllAWSWAFRegionalRule_Predicate(template *Template) map[string]*AWSWAFRegionalRule_Predicate {

	results := map[string]*AWSWAFRegionalRule_Predicate{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalRule_Predicate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalRule_PredicateWithName retrieves all AWSWAFRegionalRule_Predicate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalRule_Predicate(name string, template *Template) (*AWSWAFRegionalRule_Predicate, error) {

	result := &AWSWAFRegionalRule_Predicate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalRule_Predicate{}, errors.New("resource not found")

}
