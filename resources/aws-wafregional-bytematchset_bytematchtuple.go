package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::ByteMatchSet.ByteMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html
type AWSWAFRegionalByteMatchSet_ByteMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-fieldtomatch
	FieldToMatch AWSWAFRegionalByteMatchSet_FieldToMatch `json:"FieldToMatch"`

	// PositionalConstraint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-positionalconstraint
	PositionalConstraint string `json:"PositionalConstraint"`

	// TargetString AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-targetstring
	TargetString string `json:"TargetString"`

	// TargetStringBase64 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-targetstringbase64
	TargetStringBase64 string `json:"TargetStringBase64"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-bytematchtuple.html#cfn-wafregional-bytematchset-bytematchtuple-texttransformation
	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalByteMatchSet_ByteMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::ByteMatchSet.ByteMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalByteMatchSet_ByteMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalByteMatchSet_ByteMatchTupleResources retrieves all AWSWAFRegionalByteMatchSet_ByteMatchTuple items from a CloudFormation template
func GetAllAWSWAFRegionalByteMatchSet_ByteMatchTuple(template *Template) map[string]*AWSWAFRegionalByteMatchSet_ByteMatchTuple {

	results := map[string]*AWSWAFRegionalByteMatchSet_ByteMatchTuple{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalByteMatchSet_ByteMatchTuple{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalByteMatchSet_ByteMatchTupleWithName retrieves all AWSWAFRegionalByteMatchSet_ByteMatchTuple items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalByteMatchSet_ByteMatchTuple(name string, template *Template) (*AWSWAFRegionalByteMatchSet_ByteMatchTuple, error) {

	result := &AWSWAFRegionalByteMatchSet_ByteMatchTuple{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalByteMatchSet_ByteMatchTuple{}, errors.New("resource not found")

}
