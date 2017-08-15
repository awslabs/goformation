package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.EbsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
type AWSEMRInstanceGroupConfig_EbsConfiguration struct {

	// EbsBlockDeviceConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs []AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_EbsConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.EbsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_EbsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_EbsConfigurationResources retrieves all AWSEMRInstanceGroupConfig_EbsConfiguration items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_EbsConfiguration(template *Template) map[string]*AWSEMRInstanceGroupConfig_EbsConfiguration {

	results := map[string]*AWSEMRInstanceGroupConfig_EbsConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_EbsConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_EbsConfigurationWithName retrieves all AWSEMRInstanceGroupConfig_EbsConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_EbsConfiguration(name string, template *Template) (*AWSEMRInstanceGroupConfig_EbsConfiguration, error) {

	result := &AWSEMRInstanceGroupConfig_EbsConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_EbsConfiguration{}, errors.New("resource not found")

}
