package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::SqlInjectionMatchSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html
type AWSWAFRegionalSqlInjectionMatchSet struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-name
	Name string `json:"Name"`

	// SqlInjectionMatchTuples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuples
	SqlInjectionMatchTuples []AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple `json:"SqlInjectionMatchTuples"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSqlInjectionMatchSet) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SqlInjectionMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSqlInjectionMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalSqlInjectionMatchSetResources retrieves all AWSWAFRegionalSqlInjectionMatchSet items from a CloudFormation template
func GetAllAWSWAFRegionalSqlInjectionMatchSet(template *Template) map[string]*AWSWAFRegionalSqlInjectionMatchSet {

	results := map[string]*AWSWAFRegionalSqlInjectionMatchSet{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalSqlInjectionMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalSqlInjectionMatchSetWithName retrieves all AWSWAFRegionalSqlInjectionMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalSqlInjectionMatchSet(name string, template *Template) (*AWSWAFRegionalSqlInjectionMatchSet, error) {

	result := &AWSWAFRegionalSqlInjectionMatchSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalSqlInjectionMatchSet{}, errors.New("resource not found")

}
