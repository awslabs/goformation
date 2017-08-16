package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFSizeConstraintSet AWS CloudFormation Resource (AWS::WAF::SizeConstraintSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html
type AWSWAFSizeConstraintSet struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-name
	Name string `json:"Name"`

	// SizeConstraints AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-sizeconstraints
	SizeConstraints []AWSWAFSizeConstraintSet_SizeConstraint `json:"SizeConstraints"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSizeConstraintSet) AWSCloudFormationType() string {
	return "AWS::WAF::SizeConstraintSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFSizeConstraintSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFSizeConstraintSetResources retrieves all AWSWAFSizeConstraintSet items from a CloudFormation template
func (t *Template) GetAllAWSWAFSizeConstraintSetResources() map[string]*AWSWAFSizeConstraintSet {

	results := map[string]*AWSWAFSizeConstraintSet{}
	for name, resource := range t.Resources {
		result := &AWSWAFSizeConstraintSet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFSizeConstraintSetWithName retrieves all AWSWAFSizeConstraintSet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFSizeConstraintSetWithName(name string) (*AWSWAFSizeConstraintSet, error) {

	result := &AWSWAFSizeConstraintSet{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFSizeConstraintSet{}, errors.New("resource not found")

}
