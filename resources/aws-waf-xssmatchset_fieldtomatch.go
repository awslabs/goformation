package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::XssMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html
type AWSWAFXssMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-data

	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::XssMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFXssMatchSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFXssMatchSet_FieldToMatchResources retrieves all AWSWAFXssMatchSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFXssMatchSet_FieldToMatch(template *Template) map[string]*AWSWAFXssMatchSet_FieldToMatch {

	results := map[string]*AWSWAFXssMatchSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFXssMatchSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFXssMatchSet_FieldToMatchWithName retrieves all AWSWAFXssMatchSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFXssMatchSet_FieldToMatch(name string, template *Template) (*AWSWAFXssMatchSet_FieldToMatch, error) {

	result := &AWSWAFXssMatchSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFXssMatchSet_FieldToMatch{}, errors.New("resource not found")

}
