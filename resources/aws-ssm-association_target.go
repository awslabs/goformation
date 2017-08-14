package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SSM::Association.Target AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html
type AWSSSMAssociation_Target struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-key

	Key string `json:"Key"`

	// Values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-target.html#cfn-ssm-association-target-values

	Values []string `json:"Values"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation_Target) AWSCloudFormationType() string {
	return "AWS::SSM::Association.Target"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMAssociation_Target) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSSMAssociation_TargetResources retrieves all AWSSSMAssociation_Target items from a CloudFormation template
func GetAllAWSSSMAssociation_Target(template *Template) map[string]*AWSSSMAssociation_Target {

	results := map[string]*AWSSSMAssociation_Target{}
	for name, resource := range template.Resources {
		result := &AWSSSMAssociation_Target{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSSMAssociation_TargetWithName retrieves all AWSSSMAssociation_Target items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSSMAssociation_Target(name string, template *Template) (*AWSSSMAssociation_Target, error) {

	result := &AWSSSMAssociation_Target{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSSMAssociation_Target{}, errors.New("resource not found")

}
