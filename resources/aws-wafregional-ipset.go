package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFRegionalIPSet AWS CloudFormation Resource (AWS::WAFRegional::IPSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html
type AWSWAFRegionalIPSet struct {

	// IPSetDescriptors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html#cfn-wafregional-ipset-ipsetdescriptors
	IPSetDescriptors []AWSWAFRegionalIPSet_IPSetDescriptor `json:"IPSetDescriptors"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html#cfn-wafregional-ipset-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalIPSet) AWSCloudFormationType() string {
	return "AWS::WAFRegional::IPSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalIPSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalIPSetResources retrieves all AWSWAFRegionalIPSet items from a CloudFormation template
func GetAllAWSWAFRegionalIPSetResources(template *Template) map[string]*AWSWAFRegionalIPSet {

	results := map[string]*AWSWAFRegionalIPSet{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalIPSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalIPSetWithName retrieves all AWSWAFRegionalIPSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSWAFRegionalIPSetWithName(name string, template *Template) (*AWSWAFRegionalIPSet, error) {

	result := &AWSWAFRegionalIPSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalIPSet{}, errors.New("resource not found")

}
