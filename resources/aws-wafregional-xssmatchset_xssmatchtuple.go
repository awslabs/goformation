package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::XssMatchSet.XssMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html
type AWSWAFRegionalXssMatchSet_XssMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-fieldtomatch

	FieldToMatch AWSWAFRegionalXssMatchSet_FieldToMatch `json:"FieldToMatch"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::XssMatchSet.XssMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalXssMatchSet_XssMatchTupleResources retrieves all AWSWAFRegionalXssMatchSet_XssMatchTuple items from a CloudFormation template
func GetAllAWSWAFRegionalXssMatchSet_XssMatchTuple(template *Template) map[string]*AWSWAFRegionalXssMatchSet_XssMatchTuple {

	results := map[string]*AWSWAFRegionalXssMatchSet_XssMatchTuple{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalXssMatchSet_XssMatchTuple{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalXssMatchSet_XssMatchTupleWithName retrieves all AWSWAFRegionalXssMatchSet_XssMatchTuple items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalXssMatchSet_XssMatchTuple(name string, template *Template) (*AWSWAFRegionalXssMatchSet_XssMatchTuple, error) {

	result := &AWSWAFRegionalXssMatchSet_XssMatchTuple{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalXssMatchSet_XssMatchTuple{}, errors.New("resource not found")

}
