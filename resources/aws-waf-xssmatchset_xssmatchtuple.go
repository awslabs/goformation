package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::XssMatchSet.XssMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html
type AWSWAFXssMatchSet_XssMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch

	FieldToMatch AWSWAFXssMatchSet_FieldToMatch `json:"FieldToMatch"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet_XssMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAF::XssMatchSet.XssMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFXssMatchSet_XssMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFXssMatchSet_XssMatchTupleResources retrieves all AWSWAFXssMatchSet_XssMatchTuple items from a CloudFormation template
func GetAllAWSWAFXssMatchSet_XssMatchTuple(template *Template) map[string]*AWSWAFXssMatchSet_XssMatchTuple {

	results := map[string]*AWSWAFXssMatchSet_XssMatchTuple{}
	for name, resource := range template.Resources {
		result := &AWSWAFXssMatchSet_XssMatchTuple{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFXssMatchSet_XssMatchTupleWithName retrieves all AWSWAFXssMatchSet_XssMatchTuple items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFXssMatchSet_XssMatchTuple(name string, template *Template) (*AWSWAFXssMatchSet_XssMatchTuple, error) {

	result := &AWSWAFXssMatchSet_XssMatchTuple{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFXssMatchSet_XssMatchTuple{}, errors.New("resource not found")

}
