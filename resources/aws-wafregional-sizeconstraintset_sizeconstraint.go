package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::SizeConstraintSet.SizeConstraint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html
type AWSWAFRegionalSizeConstraintSet_SizeConstraint struct {

	// ComparisonOperator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-comparisonoperator

	ComparisonOperator string `json:"ComparisonOperator"`

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-fieldtomatch

	FieldToMatch AWSWAFRegionalSizeConstraintSet_FieldToMatch `json:"FieldToMatch"`

	// Size AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-size

	Size int64 `json:"Size"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-sizeconstraint.html#cfn-wafregional-sizeconstraintset-sizeconstraint-texttransformation

	TextTransformation string `json:"TextTransformation"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSizeConstraintSet_SizeConstraint) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SizeConstraintSet.SizeConstraint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSizeConstraintSet_SizeConstraint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalSizeConstraintSet_SizeConstraintResources retrieves all AWSWAFRegionalSizeConstraintSet_SizeConstraint items from a CloudFormation template
func GetAllAWSWAFRegionalSizeConstraintSet_SizeConstraint(template *Template) map[string]*AWSWAFRegionalSizeConstraintSet_SizeConstraint {

	results := map[string]*AWSWAFRegionalSizeConstraintSet_SizeConstraint{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalSizeConstraintSet_SizeConstraint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalSizeConstraintSet_SizeConstraintWithName retrieves all AWSWAFRegionalSizeConstraintSet_SizeConstraint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalSizeConstraintSet_SizeConstraint(name string, template *Template) (*AWSWAFRegionalSizeConstraintSet_SizeConstraint, error) {

	result := &AWSWAFRegionalSizeConstraintSet_SizeConstraint{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalSizeConstraintSet_SizeConstraint{}, errors.New("resource not found")

}
