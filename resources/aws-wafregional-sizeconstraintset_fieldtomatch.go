package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::SizeConstraintSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html
type AWSWAFRegionalSizeConstraintSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html#cfn-wafregional-sizeconstraintset-fieldtomatch-data

	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sizeconstraintset-fieldtomatch.html#cfn-wafregional-sizeconstraintset-fieldtomatch-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSizeConstraintSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SizeConstraintSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSizeConstraintSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalSizeConstraintSet_FieldToMatchResources retrieves all AWSWAFRegionalSizeConstraintSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFRegionalSizeConstraintSet_FieldToMatch(template *Template) map[string]*AWSWAFRegionalSizeConstraintSet_FieldToMatch {

	results := map[string]*AWSWAFRegionalSizeConstraintSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalSizeConstraintSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalSizeConstraintSet_FieldToMatchWithName retrieves all AWSWAFRegionalSizeConstraintSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalSizeConstraintSet_FieldToMatch(name string, template *Template) (*AWSWAFRegionalSizeConstraintSet_FieldToMatch, error) {

	result := &AWSWAFRegionalSizeConstraintSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalSizeConstraintSet_FieldToMatch{}, errors.New("resource not found")

}
