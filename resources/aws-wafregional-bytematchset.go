package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFRegionalByteMatchSet AWS CloudFormation Resource (AWS::WAFRegional::ByteMatchSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html
type AWSWAFRegionalByteMatchSet struct {

	// ByteMatchTuples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-bytematchtuples
	ByteMatchTuples []AWSWAFRegionalByteMatchSet_ByteMatchTuple `json:"ByteMatchTuples"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalByteMatchSet) AWSCloudFormationType() string {
	return "AWS::WAFRegional::ByteMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalByteMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalByteMatchSetResources retrieves all AWSWAFRegionalByteMatchSet items from a CloudFormation template
func GetAllAWSWAFRegionalByteMatchSetResources(template *Template) map[string]*AWSWAFRegionalByteMatchSet {

	results := map[string]*AWSWAFRegionalByteMatchSet{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalByteMatchSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalByteMatchSetWithName retrieves all AWSWAFRegionalByteMatchSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSWAFRegionalByteMatchSetWithName(name string, template *Template) (*AWSWAFRegionalByteMatchSet, error) {

	result := &AWSWAFRegionalByteMatchSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalByteMatchSet{}, errors.New("resource not found")

}
