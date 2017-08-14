package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::SizeConstraintSet.FieldToMatch AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html
type AWSWAFSizeConstraintSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data

	Data string `json:"Data"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSizeConstraintSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::SizeConstraintSet.FieldToMatch"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSizeConstraintSet_FieldToMatch) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFSizeConstraintSet_FieldToMatchResources retrieves all AWSWAFSizeConstraintSet_FieldToMatch items from a CloudFormation template
func GetAllAWSWAFSizeConstraintSet_FieldToMatch(template *Template) map[string]*AWSWAFSizeConstraintSet_FieldToMatch {

	results := map[string]*AWSWAFSizeConstraintSet_FieldToMatch{}
	for name, resource := range template.Resources {
		result := &AWSWAFSizeConstraintSet_FieldToMatch{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSizeConstraintSet_FieldToMatchWithName retrieves all AWSWAFSizeConstraintSet_FieldToMatch items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFSizeConstraintSet_FieldToMatch(name string, template *Template) (*AWSWAFSizeConstraintSet_FieldToMatch, error) {

	result := &AWSWAFSizeConstraintSet_FieldToMatch{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSizeConstraintSet_FieldToMatch{}, errors.New("resource not found")

}
