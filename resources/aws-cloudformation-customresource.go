package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCloudFormationCustomResource AWS CloudFormation Resource (AWS::CloudFormation::CustomResource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
type AWSCloudFormationCustomResource struct {

	// ServiceToken AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html#cfn-customresource-servicetoken
	ServiceToken string `json:"ServiceToken"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationCustomResource) AWSCloudFormationType() string {
	return "AWS::CloudFormation::CustomResource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationCustomResource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFormationCustomResourceResources retrieves all AWSCloudFormationCustomResource items from a CloudFormation template
func (t *Template) GetAllAWSCloudFormationCustomResourceResources() map[string]*AWSCloudFormationCustomResource {

	results := map[string]*AWSCloudFormationCustomResource{}
	for name, resource := range t.Resources {
		result := &AWSCloudFormationCustomResource{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFormationCustomResourceWithName retrieves all AWSCloudFormationCustomResource items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationCustomResourceWithName(name string) (*AWSCloudFormationCustomResource, error) {

	result := &AWSCloudFormationCustomResource{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFormationCustomResource{}, errors.New("resource not found")

}
