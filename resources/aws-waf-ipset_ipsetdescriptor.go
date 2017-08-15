package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::IPSet.IPSetDescriptor AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html
type AWSWAFIPSet_IPSetDescriptor struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html#cfn-waf-ipset-ipsetdescriptors-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-ipset-ipsetdescriptors.html#cfn-waf-ipset-ipsetdescriptors-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFIPSet_IPSetDescriptor) AWSCloudFormationType() string {
	return "AWS::WAF::IPSet.IPSetDescriptor"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFIPSet_IPSetDescriptor) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFIPSet_IPSetDescriptorResources retrieves all AWSWAFIPSet_IPSetDescriptor items from a CloudFormation template
func GetAllAWSWAFIPSet_IPSetDescriptor(template *Template) map[string]*AWSWAFIPSet_IPSetDescriptor {

	results := map[string]*AWSWAFIPSet_IPSetDescriptor{}
	for name, resource := range template.Resources {
		result := &AWSWAFIPSet_IPSetDescriptor{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFIPSet_IPSetDescriptorWithName retrieves all AWSWAFIPSet_IPSetDescriptor items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFIPSet_IPSetDescriptor(name string, template *Template) (*AWSWAFIPSet_IPSetDescriptor, error) {

	result := &AWSWAFIPSet_IPSetDescriptor{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFIPSet_IPSetDescriptor{}, errors.New("resource not found")

}
