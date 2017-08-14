package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::SizeConstraintSet.SizeConstraint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html
type AWSWAFSizeConstraintSet_SizeConstraint struct {

	// ComparisonOperator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-comparisonoperator

	ComparisonOperator string `json:"ComparisonOperator"`

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch

	FieldToMatch AWSWAFSizeConstraintSet_FieldToMatch `json:"FieldToMatch"`

	// Size AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-size

	Size int64 `json:"Size"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html#cfn-waf-sizeconstraintset-sizeconstraint-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSizeConstraintSet_SizeConstraint) AWSCloudFormationType() string {
	return "AWS::WAF::SizeConstraintSet.SizeConstraint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSizeConstraintSet_SizeConstraint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFSizeConstraintSet_SizeConstraintResources retrieves all AWSWAFSizeConstraintSet_SizeConstraint items from a CloudFormation template
func GetAllAWSWAFSizeConstraintSet_SizeConstraint(template *Template) map[string]*AWSWAFSizeConstraintSet_SizeConstraint {

	results := map[string]*AWSWAFSizeConstraintSet_SizeConstraint{}
	for name, resource := range template.Resources {
		result := &AWSWAFSizeConstraintSet_SizeConstraint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSizeConstraintSet_SizeConstraintWithName retrieves all AWSWAFSizeConstraintSet_SizeConstraint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFSizeConstraintSet_SizeConstraint(name string, template *Template) (*AWSWAFSizeConstraintSet_SizeConstraint, error) {

	result := &AWSWAFSizeConstraintSet_SizeConstraint{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSizeConstraintSet_SizeConstraint{}, errors.New("resource not found")

}
