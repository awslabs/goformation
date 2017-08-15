package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticBeanstalk::Environment.OptionSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type AWSElasticBeanstalkEnvironment_OptionSettings struct {

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
func (r *AWSElasticBeanstalkEnvironment_OptionSettings) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Environment.OptionSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkEnvironment_OptionSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkEnvironment_OptionSettingsResources retrieves all AWSElasticBeanstalkEnvironment_OptionSettings items from a CloudFormation template
func GetAllAWSElasticBeanstalkEnvironment_OptionSettings(template *Template) map[string]*AWSElasticBeanstalkEnvironment_OptionSettings {

	results := map[string]*AWSElasticBeanstalkEnvironment_OptionSettings{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkEnvironment_OptionSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkEnvironment_OptionSettingsWithName retrieves all AWSElasticBeanstalkEnvironment_OptionSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticBeanstalkEnvironment_OptionSettings(name string, template *Template) (*AWSElasticBeanstalkEnvironment_OptionSettings, error) {

	result := &AWSElasticBeanstalkEnvironment_OptionSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkEnvironment_OptionSettings{}, errors.New("resource not found")

}
