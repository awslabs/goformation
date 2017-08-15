package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.VolumeSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html
type AWSEMRCluster_VolumeSpecification struct {

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
func (r *AWSEMRCluster_VolumeSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.VolumeSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_VolumeSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_VolumeSpecificationResources retrieves all AWSEMRCluster_VolumeSpecification items from a CloudFormation template
func GetAllAWSEMRCluster_VolumeSpecification(template *Template) map[string]*AWSEMRCluster_VolumeSpecification {

	results := map[string]*AWSEMRCluster_VolumeSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_VolumeSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_VolumeSpecificationWithName retrieves all AWSEMRCluster_VolumeSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_VolumeSpecification(name string, template *Template) (*AWSEMRCluster_VolumeSpecification, error) {

	result := &AWSEMRCluster_VolumeSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_VolumeSpecification{}, errors.New("resource not found")

}
