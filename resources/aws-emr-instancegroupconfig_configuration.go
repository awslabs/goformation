package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.Configuration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type AWSEMRInstanceGroupConfig_Configuration struct {

	// Classification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-classification

	Classification string `json:"Classification"`

	// ConfigurationProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurationproperties

	ConfigurationProperties map[string]string `json:"ConfigurationProperties"`

	// Configurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurations

	Configurations []AWSEMRInstanceGroupConfig_Configuration `json:"Configurations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_Configuration) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.Configuration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_Configuration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_ConfigurationResources retrieves all AWSEMRInstanceGroupConfig_Configuration items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_Configuration(template *Template) map[string]*AWSEMRInstanceGroupConfig_Configuration {

	results := map[string]*AWSEMRInstanceGroupConfig_Configuration{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_Configuration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_ConfigurationWithName retrieves all AWSEMRInstanceGroupConfig_Configuration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_Configuration(name string, template *Template) (*AWSEMRInstanceGroupConfig_Configuration, error) {

	result := &AWSEMRInstanceGroupConfig_Configuration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_Configuration{}, errors.New("resource not found")

}
