package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSSDBDomain AWS CloudFormation Resource (AWS::SDB::Domain)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html
type AWSSDBDomain struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html#cfn-sdb-domain-description
	Description string `json:"Description"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSDBDomain) AWSCloudFormationType() string {
	return "AWS::SDB::Domain"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSDBDomain) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSDBDomainResources retrieves all AWSSDBDomain items from a CloudFormation template
func (t *Template) GetAllAWSSDBDomainResources() map[string]*AWSSDBDomain {

	results := map[string]*AWSSDBDomain{}
	for name, resource := range t.Resources {
		result := &AWSSDBDomain{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSDBDomainWithName retrieves all AWSSDBDomain items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSDBDomainWithName(name string) (*AWSSDBDomain, error) {

	result := &AWSSDBDomain{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSDBDomain{}, errors.New("resource not found")

}
