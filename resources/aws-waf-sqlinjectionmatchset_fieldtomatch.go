package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::SqlInjectionMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html
type AWSWAFSqlInjectionMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data
	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSqlInjectionMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::SqlInjectionMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSqlInjectionMatchSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFSqlInjectionMatchSet_FieldToMatchResources retrieves all AWSWAFSqlInjectionMatchSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFSqlInjectionMatchSet_FieldToMatch(template *Template) map[string]*AWSWAFSqlInjectionMatchSet_FieldToMatch {

	results := map[string]*AWSWAFSqlInjectionMatchSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFSqlInjectionMatchSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSqlInjectionMatchSet_FieldToMatchWithName retrieves all AWSWAFSqlInjectionMatchSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFSqlInjectionMatchSet_FieldToMatch(name string, template *Template) (*AWSWAFSqlInjectionMatchSet_FieldToMatch, error) {

	result := &AWSWAFSqlInjectionMatchSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSqlInjectionMatchSet_FieldToMatch{}, errors.New("resource not found")

}
