package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.EbsBlockDeviceConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html
type AWSEMRCluster_EbsBlockDeviceConfig struct {

	// VolumeSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification

	VolumeSpecification AWSEMRCluster_VolumeSpecification `json:"VolumeSpecification"`

	// VolumesPerInstance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumesperinstance

	VolumesPerInstance int64 `json:"VolumesPerInstance"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_EbsBlockDeviceConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.EbsBlockDeviceConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_EbsBlockDeviceConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_EbsBlockDeviceConfigResources retrieves all AWSEMRCluster_EbsBlockDeviceConfig items from a CloudFormation template
func GetAllAWSEMRCluster_EbsBlockDeviceConfig(template *Template) map[string]*AWSEMRCluster_EbsBlockDeviceConfig {

	results := map[string]*AWSEMRCluster_EbsBlockDeviceConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_EbsBlockDeviceConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_EbsBlockDeviceConfigWithName retrieves all AWSEMRCluster_EbsBlockDeviceConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_EbsBlockDeviceConfig(name string, template *Template) (*AWSEMRCluster_EbsBlockDeviceConfig, error) {

	result := &AWSEMRCluster_EbsBlockDeviceConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_EbsBlockDeviceConfig{}, errors.New("resource not found")

}
