package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::SqlInjectionMatchSet AWS CloudFormation Resource
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
func GetAllAWSWAFSqlInjectionMatchSet(template *Template) map[string]*AWSWAFSqlInjectionMatchSet {

	results := map[string]*AWSWAFSqlInjectionMatchSet{}
	for name, resource := range template.Resources {
		result := &AWSWAFSqlInjectionMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSqlInjectionMatchSetWithName retrieves all AWSWAFSqlInjectionMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFSqlInjectionMatchSet(name string, template *Template) (*AWSWAFSqlInjectionMatchSet, error) {

	result := &AWSWAFSqlInjectionMatchSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSqlInjectionMatchSet{}, errors.New("resource not found")

}
