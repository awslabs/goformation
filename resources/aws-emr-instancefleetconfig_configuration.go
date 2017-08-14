package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.Configuration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-configuration.html
type AWSEMRInstanceFleetConfig_Configuration struct {

	// Classification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-configuration.html#cfn-elasticmapreduce-instancefleetconfig-configuration-classification

	Classification string `json:"Classification"`

	// ConfigurationProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-configuration.html#cfn-elasticmapreduce-instancefleetconfig-configuration-configurationproperties

	ConfigurationProperties map[string]string `json:"ConfigurationProperties"`

	// Configurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-configuration.html#cfn-elasticmapreduce-instancefleetconfig-configuration-configurations

	Configurations []AWSEMRInstanceFleetConfig_Configuration `json:"Configurations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_Configuration) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.Configuration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_Configuration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_ConfigurationResources retrieves all AWSEMRInstanceFleetConfig_Configuration items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_Configuration(template *Template) map[string]*AWSEMRInstanceFleetConfig_Configuration {

	results := map[string]*AWSEMRInstanceFleetConfig_Configuration{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_Configuration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_ConfigurationWithName retrieves all AWSEMRInstanceFleetConfig_Configuration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_Configuration(name string, template *Template) (*AWSEMRInstanceFleetConfig_Configuration, error) {

	result := &AWSEMRInstanceFleetConfig_Configuration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_Configuration{}, errors.New("resource not found")

}
