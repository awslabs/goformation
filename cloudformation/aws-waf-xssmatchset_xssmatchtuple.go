package cloudformation

import (
	"encoding/json"
)

// AWSWAFXssMatchSet_XssMatchTuple AWS CloudFormation Resource (AWS::WAF::XssMatchSet.XssMatchTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html
type AWSWAFXssMatchSet_XssMatchTuple struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch
	FieldToMatch *AWSWAFXssMatchSet_FieldToMatch `json:"FieldToMatch,omitempty"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple.html#cfn-waf-xssmatchset-xssmatchtuple-texttransformation
	TextTransformation *Value `json:"TextTransformation,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet_XssMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAF::XssMatchSet.XssMatchTuple"
}

func (r *AWSWAFXssMatchSet_XssMatchTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
