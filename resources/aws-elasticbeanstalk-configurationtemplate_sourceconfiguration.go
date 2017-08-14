package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html
type AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html#cfn-beanstalk-configurationtemplate-sourceconfiguration-applicationname

	ApplicationName string `json:"ApplicationName"`

	// TemplateName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-configurationtemplate-sourceconfiguration.html#cfn-beanstalk-configurationtemplate-sourceconfiguration-templatename

	TemplateName string `json:"TemplateName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate.SourceConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkConfigurationTemplate_SourceConfigurationResources retrieves all AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration items from a CloudFormation template
func GetAllAWSElasticBeanstalkConfigurationTemplate_SourceConfiguration(template *Template) map[string]*AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration {

	results := map[string]*AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkConfigurationTemplate_SourceConfigurationWithName retrieves all AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticBeanstalkConfigurationTemplate_SourceConfiguration(name string, template *Template) (*AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration, error) {

	result := &AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkConfigurationTemplate_SourceConfiguration{}, errors.New("resource not found")

}
