package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Volume AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html
type AWSOpsWorksVolume struct {

	// Ec2VolumeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-ec2volumeid
	Ec2VolumeId string `json:"Ec2VolumeId"`

	// MountPoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-mountpoint
	MountPoint string `json:"MountPoint"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-name
	Name string `json:"Name"`

	// StackId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html#cfn-opsworks-volume-stackid
	StackId string `json:"StackId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksVolume) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Volume"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksVolume) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksVolumeResources retrieves all AWSOpsWorksVolume items from a CloudFormation template
func GetAllAWSOpsWorksVolumeResources(template *Template) map[string]*AWSOpsWorksVolume {

	results := map[string]*AWSOpsWorksVolume{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksVolume{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksVolumeWithName retrieves all AWSOpsWorksVolume items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSOpsWorksVolumeWithName(name string, template *Template) (*AWSOpsWorksVolume, error) {

	result := &AWSOpsWorksVolume{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksVolume{}, errors.New("resource not found")

}
