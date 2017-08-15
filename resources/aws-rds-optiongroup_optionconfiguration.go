package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::OptionGroup.OptionConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html
type AWSRDSOptionGroup_OptionConfiguration struct {

	// DBSecurityGroupMemberships AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-dbsecuritygroupmemberships
	DBSecurityGroupMemberships []string `json:"DBSecurityGroupMemberships"`

	// OptionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-optionname
	OptionName string `json:"OptionName"`

	// OptionSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-optionsettings
	OptionSettings AWSRDSOptionGroup_OptionSetting `json:"OptionSettings"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-port
	Port int64 `json:"Port"`

	// VpcSecurityGroupMemberships AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfigurations.html#cfn-rds-optiongroup-optionconfigurations-vpcsecuritygroupmemberships
	VpcSecurityGroupMemberships []string `json:"VpcSecurityGroupMemberships"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSOptionGroup_OptionConfiguration) AWSCloudFormationType() string {
	return "AWS::RDS::OptionGroup.OptionConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSOptionGroup_OptionConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSOptionGroup_OptionConfigurationResources retrieves all AWSRDSOptionGroup_OptionConfiguration items from a CloudFormation template
func GetAllAWSRDSOptionGroup_OptionConfiguration(template *Template) map[string]*AWSRDSOptionGroup_OptionConfiguration {

	results := map[string]*AWSRDSOptionGroup_OptionConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSRDSOptionGroup_OptionConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSOptionGroup_OptionConfigurationWithName retrieves all AWSRDSOptionGroup_OptionConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRDSOptionGroup_OptionConfiguration(name string, template *Template) (*AWSRDSOptionGroup_OptionConfiguration, error) {

	result := &AWSRDSOptionGroup_OptionConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSOptionGroup_OptionConfiguration{}, errors.New("resource not found")

}
