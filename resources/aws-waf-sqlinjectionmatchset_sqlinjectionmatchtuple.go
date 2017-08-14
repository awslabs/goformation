package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html
type AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-fieldtomatch

	FieldToMatch AWSWAFSqlInjectionMatchSet_FieldToMatch `json:"FieldToMatch"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFSqlInjectionMatchSet_SqlInjectionMatchTupleResources retrieves all AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple items from a CloudFormation template
func GetAllAWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple(template *Template) map[string]*AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple {

	results := map[string]*AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple{}
	for name, resource := range template.Resources {
		result := &AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSqlInjectionMatchSet_SqlInjectionMatchTupleWithName retrieves all AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple(name string, template *Template) (*AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple, error) {

	result := &AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple{}, errors.New("resource not found")

}
