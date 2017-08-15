package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::XssMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html
type AWSWAFRegionalXssMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-data
	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAFRegional::XssMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalXssMatchSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalXssMatchSet_FieldToMatchResources retrieves all AWSWAFRegionalXssMatchSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFRegionalXssMatchSet_FieldToMatch(template *Template) map[string]*AWSWAFRegionalXssMatchSet_FieldToMatch {

	results := map[string]*AWSWAFRegionalXssMatchSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalXssMatchSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalXssMatchSet_FieldToMatchWithName retrieves all AWSWAFRegionalXssMatchSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalXssMatchSet_FieldToMatch(name string, template *Template) (*AWSWAFRegionalXssMatchSet_FieldToMatch, error) {

	result := &AWSWAFRegionalXssMatchSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalXssMatchSet_FieldToMatch{}, errors.New("resource not found")

}
