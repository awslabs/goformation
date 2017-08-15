package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SSM::Association.ParameterValues AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html
type AWSSSMAssociation_ParameterValues struct {

	// ParameterValues AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
	ParameterValues []string `json:"ParameterValues"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation_ParameterValues) AWSCloudFormationType() string {
	return "AWS::SSM::Association.ParameterValues"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMAssociation_ParameterValues) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSSMAssociation_ParameterValuesResources retrieves all AWSSSMAssociation_ParameterValues items from a CloudFormation template
func GetAllAWSSSMAssociation_ParameterValues(template *Template) map[string]*AWSSSMAssociation_ParameterValues {

	results := map[string]*AWSSSMAssociation_ParameterValues{}
	for name, resource := range template.Resources {
		result := &AWSSSMAssociation_ParameterValues{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSSMAssociation_ParameterValuesWithName retrieves all AWSSSMAssociation_ParameterValues items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSSMAssociation_ParameterValues(name string, template *Template) (*AWSSSMAssociation_ParameterValues, error) {

	result := &AWSSSMAssociation_ParameterValues{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSSMAssociation_ParameterValues{}, errors.New("resource not found")

}
