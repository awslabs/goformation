package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::ByteMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html
type AWSWAFRegionalByteMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-data

	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalByteMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAFRegional::ByteMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalByteMatchSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalByteMatchSet_FieldToMatchResources retrieves all AWSWAFRegionalByteMatchSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFRegionalByteMatchSet_FieldToMatch(template *Template) map[string]*AWSWAFRegionalByteMatchSet_FieldToMatch {

	results := map[string]*AWSWAFRegionalByteMatchSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalByteMatchSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalByteMatchSet_FieldToMatchWithName retrieves all AWSWAFRegionalByteMatchSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalByteMatchSet_FieldToMatch(name string, template *Template) (*AWSWAFRegionalByteMatchSet_FieldToMatch, error) {

	result := &AWSWAFRegionalByteMatchSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalByteMatchSet_FieldToMatch{}, errors.New("resource not found")

}
