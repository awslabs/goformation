package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.EbsBlockDeviceConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html
type AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig struct {

	// VolumeSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html#cfn-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig-volumespecification
	VolumeSpecification AWSEMRInstanceFleetConfig_VolumeSpecification `json:"VolumeSpecification"`

	// VolumesPerInstance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html#cfn-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig-volumesperinstance
	VolumesPerInstance int64 `json:"VolumesPerInstance"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.EbsBlockDeviceConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_EbsBlockDeviceConfigResources retrieves all AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_EbsBlockDeviceConfig(template *Template) map[string]*AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig {

	results := map[string]*AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_EbsBlockDeviceConfigWithName retrieves all AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_EbsBlockDeviceConfig(name string, template *Template) (*AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig, error) {

	result := &AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig{}, errors.New("resource not found")

}
