package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// Tag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationinstance-tag.html
type Tag struct {

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationinstance-tag.html#cfn-dms-replicationinstance-tag-key

	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationinstance-tag.html#cfn-dms-replicationinstance-tag-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Tag) AWSCloudFormationType() string {
	return "Tag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *Tag) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllTagResources retrieves all Tag items from a CloudFormation template
func GetAllTag(template *Template) map[string]*Tag {

	results := map[string]*Tag{}
	for name, resource := range template.Resources {
		result := &Tag{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetTagWithName retrieves all Tag items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameTag(name string, template *Template) (*Tag, error) {

	result := &Tag{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &Tag{}, errors.New("resource not found")

}
