package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html
type AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-fieldtomatch
	FieldToMatch AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch `json:"FieldToMatch"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-texttransformation
	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTupleResources retrieves all AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple items from a CloudFormation template
func GetAllAWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple(template *Template) map[string]*AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple {

	results := map[string]*AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTupleWithName retrieves all AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple(name string, template *Template) (*AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple, error) {

	result := &AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple{}, errors.New("resource not found")

}
