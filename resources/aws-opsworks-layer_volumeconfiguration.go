package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.VolumeConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html
type AWSOpsWorksLayer_VolumeConfiguration struct {

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-iops
	Iops int64 `json:"Iops"`

	// MountPoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-mountpoint
	MountPoint string `json:"MountPoint"`

	// NumberOfDisks AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-numberofdisks
	NumberOfDisks int64 `json:"NumberOfDisks"`

	// RaidLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-raidlevel
	RaidLevel int64 `json:"RaidLevel"`

	// Size AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-size
	Size int64 `json:"Size"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-volumeconfiguration.html#cfn-opsworks-layer-volconfig-volumetype
	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_VolumeConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.VolumeConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_VolumeConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksLayer_VolumeConfigurationResources retrieves all AWSOpsWorksLayer_VolumeConfiguration items from a CloudFormation template
func GetAllAWSOpsWorksLayer_VolumeConfiguration(template *Template) map[string]*AWSOpsWorksLayer_VolumeConfiguration {

	results := map[string]*AWSOpsWorksLayer_VolumeConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksLayer_VolumeConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksLayer_VolumeConfigurationWithName retrieves all AWSOpsWorksLayer_VolumeConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksLayer_VolumeConfiguration(name string, template *Template) (*AWSOpsWorksLayer_VolumeConfiguration, error) {

	result := &AWSOpsWorksLayer_VolumeConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksLayer_VolumeConfiguration{}, errors.New("resource not found")

}
