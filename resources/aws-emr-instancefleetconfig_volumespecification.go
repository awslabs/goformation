package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceFleetConfig.VolumeSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html
type AWSEMRInstanceFleetConfig_VolumeSpecification struct {

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-iops
	Iops int64 `json:"Iops"`

	// SizeInGB AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-sizeingb
	SizeInGB int64 `json:"SizeInGB"`

	// VolumeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-volumetype
	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_VolumeSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.VolumeSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig_VolumeSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfig_VolumeSpecificationResources retrieves all AWSEMRInstanceFleetConfig_VolumeSpecification items from a CloudFormation template
func GetAllAWSEMRInstanceFleetConfig_VolumeSpecification(template *Template) map[string]*AWSEMRInstanceFleetConfig_VolumeSpecification {

	results := map[string]*AWSEMRInstanceFleetConfig_VolumeSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceFleetConfig_VolumeSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfig_VolumeSpecificationWithName retrieves all AWSEMRInstanceFleetConfig_VolumeSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceFleetConfig_VolumeSpecification(name string, template *Template) (*AWSEMRInstanceFleetConfig_VolumeSpecification, error) {

	result := &AWSEMRInstanceFleetConfig_VolumeSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig_VolumeSpecification{}, errors.New("resource not found")

}
