package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting struct {

	// Namespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-namespace
	Namespace string `json:"Namespace"`

	// OptionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-optionname
	OptionName string `json:"OptionName"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSettingResources retrieves all AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting items from a CloudFormation template
func GetAllAWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting(template *Template) map[string]*AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting {

	results := map[string]*AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSettingWithName retrieves all AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting(name string, template *Template) (*AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting, error) {

	result := &AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting{}, errors.New("resource not found")

}
