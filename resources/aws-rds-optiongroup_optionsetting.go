package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::OptionGroup.OptionSetting AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html
type AWSRDSOptionGroup_OptionSetting struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html#cfn-rds-optiongroup-optionconfigurations-optionsettings-name

	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations-optionsettings.html#cfn-rds-optiongroup-optionconfigurations-optionsettings-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSOptionGroup_OptionSetting) AWSCloudFormationType() string {
	return "AWS::RDS::OptionGroup.OptionSetting"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSOptionGroup_OptionSetting) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSOptionGroup_OptionSettingResources retrieves all AWSRDSOptionGroup_OptionSetting items from a CloudFormation template
func GetAllAWSRDSOptionGroup_OptionSetting(template *Template) map[string]*AWSRDSOptionGroup_OptionSetting {

	results := map[string]*AWSRDSOptionGroup_OptionSetting{}
	for name, resource := range template.Resources {
		result := &AWSRDSOptionGroup_OptionSetting{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSOptionGroup_OptionSettingWithName retrieves all AWSRDSOptionGroup_OptionSetting items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRDSOptionGroup_OptionSetting(name string, template *Template) (*AWSRDSOptionGroup_OptionSetting, error) {

	result := &AWSRDSOptionGroup_OptionSetting{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSOptionGroup_OptionSetting{}, errors.New("resource not found")

}
