package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFByteMatchSet AWS CloudFormation Resource (AWS::WAF::ByteMatchSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html
type AWSWAFByteMatchSet struct {

	// ByteMatchTuples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html#cfn-waf-bytematchset-bytematchtuples
	ByteMatchTuples []AWSWAFByteMatchSet_ByteMatchTuple `json:"ByteMatchTuples"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html#cfn-waf-bytematchset-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFByteMatchSet) AWSCloudFormationType() string {
	return "AWS::WAF::ByteMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFByteMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFByteMatchSetResources retrieves all AWSWAFByteMatchSet items from a CloudFormation template
func (t *Template) GetAllAWSWAFByteMatchSetResources() map[string]*AWSWAFByteMatchSet {

	results := map[string]*AWSWAFByteMatchSet{}
	for name, resource := range t.Resources {
		result := &AWSWAFByteMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFByteMatchSetWithName retrieves all AWSWAFByteMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFByteMatchSetWithName(name string) (*AWSWAFByteMatchSet, error) {

	result := &AWSWAFByteMatchSet{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFByteMatchSet{}, errors.New("resource not found")

}
