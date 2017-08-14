package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Instance.EbsBlockDevice AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html
type AWSOpsWorksInstance_EbsBlockDevice struct {

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-deleteontermination

	DeleteOnTermination bool `json:"DeleteOnTermination"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-iops

	Iops int64 `json:"Iops"`

	// SnapshotId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-snapshotid

	SnapshotId string `json:"SnapshotId"`

	// VolumeSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-volumesize

	VolumeSize int64 `json:"VolumeSize"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-volumetype

	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksInstance_EbsBlockDevice) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Instance.EbsBlockDevice"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksInstance_EbsBlockDevice) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksInstance_EbsBlockDeviceResources retrieves all AWSOpsWorksInstance_EbsBlockDevice items from a CloudFormation template
func GetAllAWSOpsWorksInstance_EbsBlockDevice(template *Template) map[string]*AWSOpsWorksInstance_EbsBlockDevice {

	results := map[string]*AWSOpsWorksInstance_EbsBlockDevice{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksInstance_EbsBlockDevice{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksInstance_EbsBlockDeviceWithName retrieves all AWSOpsWorksInstance_EbsBlockDevice items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksInstance_EbsBlockDevice(name string, template *Template) (*AWSOpsWorksInstance_EbsBlockDevice, error) {

	result := &AWSOpsWorksInstance_EbsBlockDevice{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksInstance_EbsBlockDevice{}, errors.New("resource not found")

}
