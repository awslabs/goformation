package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFSqlInjectionMatchSet AWS CloudFormation Resource (AWS::WAF::SqlInjectionMatchSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
type AWSWAFSqlInjectionMatchSet struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-name
	Name string `json:"Name"`

	// SqlInjectionMatchTuples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples
	SqlInjectionMatchTuples []AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple `json:"SqlInjectionMatchTuples"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSqlInjectionMatchSet) AWSCloudFormationType() string {
	return "AWS::WAF::SqlInjectionMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSqlInjectionMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFSqlInjectionMatchSetResources retrieves all AWSWAFSqlInjectionMatchSet items from a CloudFormation template
func (t *Template) GetAllAWSWAFSqlInjectionMatchSetResources() map[string]*AWSWAFSqlInjectionMatchSet {

	results := map[string]*AWSWAFSqlInjectionMatchSet{}
	for name, resource := range t.Resources {
		result := &AWSWAFSqlInjectionMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSqlInjectionMatchSetWithName retrieves all AWSWAFSqlInjectionMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFSqlInjectionMatchSetWithName(name string) (*AWSWAFSqlInjectionMatchSet, error) {

	result := &AWSWAFSqlInjectionMatchSet{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSqlInjectionMatchSet{}, errors.New("resource not found")

}
