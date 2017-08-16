package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRDSOptionGroup AWS CloudFormation Resource (AWS::RDS::OptionGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
type AWSRDSOptionGroup struct {

	// EngineName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-enginename
	EngineName string `json:"EngineName"`

	// MajorEngineVersion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-majorengineversion
	MajorEngineVersion string `json:"MajorEngineVersion"`

	// OptionConfigurations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optionconfigurations
	OptionConfigurations []AWSRDSOptionGroup_OptionConfiguration `json:"OptionConfigurations"`

	// OptionGroupDescription AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optiongroupdescription
	OptionGroupDescription string `json:"OptionGroupDescription"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSOptionGroup) AWSCloudFormationType() string {
	return "AWS::RDS::OptionGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSOptionGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSOptionGroupResources retrieves all AWSRDSOptionGroup items from a CloudFormation template
func (t *Template) GetAllAWSRDSOptionGroupResources() map[string]*AWSRDSOptionGroup {

	results := map[string]*AWSRDSOptionGroup{}
	for name, resource := range t.Resources {
		result := &AWSRDSOptionGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSOptionGroupWithName retrieves all AWSRDSOptionGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSOptionGroupWithName(name string) (*AWSRDSOptionGroup, error) {

	result := &AWSRDSOptionGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSOptionGroup{}, errors.New("resource not found")

}
