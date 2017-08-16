package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFXssMatchSet AWS CloudFormation Resource (AWS::WAF::XssMatchSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html
type AWSWAFXssMatchSet struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-name
	Name string `json:"Name"`

	// XssMatchTuples AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-xssmatchtuples
	XssMatchTuples []AWSWAFXssMatchSet_XssMatchTuple `json:"XssMatchTuples"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet) AWSCloudFormationType() string {
	return "AWS::WAF::XssMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFXssMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFXssMatchSetResources retrieves all AWSWAFXssMatchSet items from a CloudFormation template
func (t *Template) GetAllAWSWAFXssMatchSetResources() map[string]*AWSWAFXssMatchSet {

	results := map[string]*AWSWAFXssMatchSet{}
	for name, resource := range t.Resources {
		result := &AWSWAFXssMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFXssMatchSetWithName retrieves all AWSWAFXssMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFXssMatchSetWithName(name string) (*AWSWAFXssMatchSet, error) {

	result := &AWSWAFXssMatchSet{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFXssMatchSet{}, errors.New("resource not found")

}
