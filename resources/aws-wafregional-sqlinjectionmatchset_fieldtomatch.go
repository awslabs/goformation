package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::SqlInjectionMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-fieldtomatch.html
type AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-fieldtomatch.html#cfn-wafregional-sqlinjectionmatchset-fieldtomatch-data
	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-fieldtomatch.html#cfn-wafregional-sqlinjectionmatchset-fieldtomatch-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SqlInjectionMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalSqlInjectionMatchSet_FieldToMatchResources retrieves all AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFRegionalSqlInjectionMatchSet_FieldToMatch(template *Template) map[string]*AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch {

	results := map[string]*AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalSqlInjectionMatchSet_FieldToMatchWithName retrieves all AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalSqlInjectionMatchSet_FieldToMatch(name string, template *Template) (*AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch, error) {

	result := &AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch{}, errors.New("resource not found")

}
