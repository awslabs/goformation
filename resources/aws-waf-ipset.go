package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFIPSet AWS CloudFormation Resource (AWS::WAF::IPSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html
type AWSWAFIPSet struct {

	// IPSetDescriptors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-ipsetdescriptors
	IPSetDescriptors []AWSWAFIPSet_IPSetDescriptor `json:"IPSetDescriptors"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFIPSet) AWSCloudFormationType() string {
	return "AWS::WAF::IPSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFIPSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFIPSetResources retrieves all AWSWAFIPSet items from a CloudFormation template
func (t *Template) GetAllAWSWAFIPSetResources() map[string]*AWSWAFIPSet {

	results := map[string]*AWSWAFIPSet{}
	for name, resource := range t.Resources {
		result := &AWSWAFIPSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFIPSetWithName retrieves all AWSWAFIPSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFIPSetWithName(name string) (*AWSWAFIPSet, error) {

	result := &AWSWAFIPSet{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFIPSet{}, errors.New("resource not found")

}
