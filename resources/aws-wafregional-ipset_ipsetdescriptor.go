package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::IPSet.IPSetDescriptor AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html
type AWSWAFRegionalIPSet_IPSetDescriptor struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html#cfn-wafregional-ipset-ipsetdescriptor-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-ipset-ipsetdescriptor.html#cfn-wafregional-ipset-ipsetdescriptor-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalIPSet_IPSetDescriptor) AWSCloudFormationType() string {
	return "AWS::WAFRegional::IPSet.IPSetDescriptor"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalIPSet_IPSetDescriptor) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalIPSet_IPSetDescriptorResources retrieves all AWSWAFRegionalIPSet_IPSetDescriptor items from a CloudFormation template
func GetAllAWSWAFRegionalIPSet_IPSetDescriptor(template *Template) map[string]*AWSWAFRegionalIPSet_IPSetDescriptor {

	results := map[string]*AWSWAFRegionalIPSet_IPSetDescriptor{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalIPSet_IPSetDescriptor{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalIPSet_IPSetDescriptorWithName retrieves all AWSWAFRegionalIPSet_IPSetDescriptor items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalIPSet_IPSetDescriptor(name string, template *Template) (*AWSWAFRegionalIPSet_IPSetDescriptor, error) {

	result := &AWSWAFRegionalIPSet_IPSetDescriptor{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalIPSet_IPSetDescriptor{}, errors.New("resource not found")

}
