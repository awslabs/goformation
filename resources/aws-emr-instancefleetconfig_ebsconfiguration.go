package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.EbsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html
type AWSEMRInstanceFleetConfig_EbsConfiguration struct {

	// EbsBlockDeviceConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html#cfn-elasticmapreduce-instancefleetconfig-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs []AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html#cfn-elasticmapreduce-instancefleetconfig-ebsconfiguration-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_EbsConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.EbsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_EbsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_EbsConfigurationResources retrieves all AWSEMRInstanceFleetConfig_EbsConfiguration items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_EbsConfiguration(template *Template) map[string]*AWSEMRInstanceFleetConfig_EbsConfiguration {

	results := map[string]*AWSEMRInstanceFleetConfig_EbsConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_EbsConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_EbsConfigurationWithName retrieves all AWSEMRInstanceFleetConfig_EbsConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_EbsConfiguration(name string, template *Template) (*AWSEMRInstanceFleetConfig_EbsConfiguration, error) {

	result := &AWSEMRInstanceFleetConfig_EbsConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_EbsConfiguration{}, errors.New("resource not found")

}
