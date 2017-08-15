package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.VolumeSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html
type AWSEMRInstanceGroupConfig_VolumeSpecification struct {

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-iops
	Iops int64 `json:"Iops"`

	// SizeInGB AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-sizeingb
	SizeInGB int64 `json:"SizeInGB"`

	// VolumeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification-volumetype
	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_VolumeSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.VolumeSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_VolumeSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_VolumeSpecificationResources retrieves all AWSEMRInstanceGroupConfig_VolumeSpecification items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_VolumeSpecification(template *Template) map[string]*AWSEMRInstanceGroupConfig_VolumeSpecification {

	results := map[string]*AWSEMRInstanceGroupConfig_VolumeSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_VolumeSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_VolumeSpecificationWithName retrieves all AWSEMRInstanceGroupConfig_VolumeSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_VolumeSpecification(name string, template *Template) (*AWSEMRInstanceGroupConfig_VolumeSpecification, error) {

	result := &AWSEMRInstanceGroupConfig_VolumeSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_VolumeSpecification{}, errors.New("resource not found")

}
