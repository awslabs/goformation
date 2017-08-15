package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::XssMatchSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html
type AWSWAFRegionalXssMatchSet struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-name
	Name string `json:"Name"`

	// XssMatchTuples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-xssmatchtuples
	XssMatchTuples []AWSWAFRegionalXssMatchSet_XssMatchTuple `json:"XssMatchTuples"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSet) AWSCloudFormationType() string {
	return "AWS::WAFRegional::XssMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalXssMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalXssMatchSetResources retrieves all AWSWAFRegionalXssMatchSet items from a CloudFormation template
func GetAllAWSWAFRegionalXssMatchSet(template *Template) map[string]*AWSWAFRegionalXssMatchSet {

	results := map[string]*AWSWAFRegionalXssMatchSet{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalXssMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalXssMatchSetWithName retrieves all AWSWAFRegionalXssMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalXssMatchSet(name string, template *Template) (*AWSWAFRegionalXssMatchSet, error) {

	result := &AWSWAFRegionalXssMatchSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalXssMatchSet{}, errors.New("resource not found")

}
