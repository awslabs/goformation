package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::SizeConstraintSet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
type AWSWAFRegionalSizeConstraintSet struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-name
	Name string `json:"Name"`

	// SizeConstraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-sizeconstraints
	SizeConstraints []AWSWAFRegionalSizeConstraintSet_SizeConstraint `json:"SizeConstraints"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSizeConstraintSet) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SizeConstraintSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalSizeConstraintSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalSizeConstraintSetResources retrieves all AWSWAFRegionalSizeConstraintSet items from a CloudFormation template
func GetAllAWSWAFRegionalSizeConstraintSet(template *Template) map[string]*AWSWAFRegionalSizeConstraintSet {

	results := map[string]*AWSWAFRegionalSizeConstraintSet{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalSizeConstraintSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalSizeConstraintSetWithName retrieves all AWSWAFRegionalSizeConstraintSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalSizeConstraintSet(name string, template *Template) (*AWSWAFRegionalSizeConstraintSet, error) {

	result := &AWSWAFRegionalSizeConstraintSet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalSizeConstraintSet{}, errors.New("resource not found")

}
