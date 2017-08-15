package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
type AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig struct {

	// VolumeSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification
	VolumeSpecification AWSEMRInstanceGroupConfig_VolumeSpecification `json:"VolumeSpecification"`

	// VolumesPerInstance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance
	VolumesPerInstance int64 `json:"VolumesPerInstance"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.EbsBlockDeviceConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_EbsBlockDeviceConfigResources retrieves all AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_EbsBlockDeviceConfig(template *Template) map[string]*AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig {

	results := map[string]*AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_EbsBlockDeviceConfigWithName retrieves all AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_EbsBlockDeviceConfig(name string, template *Template) (*AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig, error) {

	result := &AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_EbsBlockDeviceConfig{}, errors.New("resource not found")

}
