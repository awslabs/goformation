package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::ByteMatchSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html
type AWSWAFByteMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch-data

	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFByteMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::ByteMatchSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFByteMatchSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFByteMatchSet_FieldToMatchResources retrieves all AWSWAFByteMatchSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFByteMatchSet_FieldToMatch(template *Template) map[string]*AWSWAFByteMatchSet_FieldToMatch {

	results := map[string]*AWSWAFByteMatchSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFByteMatchSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFByteMatchSet_FieldToMatchWithName retrieves all AWSWAFByteMatchSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFByteMatchSet_FieldToMatch(name string, template *Template) (*AWSWAFByteMatchSet_FieldToMatch, error) {

	result := &AWSWAFByteMatchSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFByteMatchSet_FieldToMatch{}, errors.New("resource not found")

}
