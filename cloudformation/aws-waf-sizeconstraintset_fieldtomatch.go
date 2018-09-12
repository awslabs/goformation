package cloudformation

import (
	"encoding/json"
)

// AWSWAFSizeConstraintSet_FieldToMatch AWS CloudFormation Resource (AWS::WAF::SizeConstraintSet.FieldToMatch)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html
type AWSWAFSizeConstraintSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data
	Data *Value `json:"Data,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSizeConstraintSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::SizeConstraintSet.FieldToMatch"
}

func (r *AWSWAFSizeConstraintSet_FieldToMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
