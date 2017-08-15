package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::ByteMatchSet.ByteMatchTuple AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html
type AWSWAFByteMatchSet_ByteMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch
	FieldToMatch AWSWAFByteMatchSet_FieldToMatch `json:"FieldToMatch"`

	// PositionalConstraint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-positionalconstraint
	PositionalConstraint string `json:"PositionalConstraint"`

	// TargetString AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstring
	TargetString string `json:"TargetString"`

	// TargetStringBase64 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstringbase64
	TargetStringBase64 string `json:"TargetStringBase64"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-texttransformation
	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFByteMatchSet_ByteMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAF::ByteMatchSet.ByteMatchTuple"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFByteMatchSet_ByteMatchTuple) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFByteMatchSet_ByteMatchTupleResources retrieves all AWSWAFByteMatchSet_ByteMatchTuple items from a CloudFormation template
func GetAllAWSWAFByteMatchSet_ByteMatchTuple(template *Template) map[string]*AWSWAFByteMatchSet_ByteMatchTuple {

	results := map[string]*AWSWAFByteMatchSet_ByteMatchTuple{}
	for name, resource := range template.Resources {
		result := &AWSWAFByteMatchSet_ByteMatchTuple{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFByteMatchSet_ByteMatchTupleWithName retrieves all AWSWAFByteMatchSet_ByteMatchTuple items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFByteMatchSet_ByteMatchTuple(name string, template *Template) (*AWSWAFByteMatchSet_ByteMatchTuple, error) {

	result := &AWSWAFByteMatchSet_ByteMatchTuple{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFByteMatchSet_ByteMatchTuple{}, errors.New("resource not found")

}
