package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.Volume AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html
type AWSEC2Instance_Volume struct {

	// Device AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html#cfn-ec2-mountpoint-device
	Device string `json:"Device"`

	// VolumeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html#cfn-ec2-mountpoint-volumeid
	VolumeId string `json:"VolumeId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_Volume) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.Volume"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_Volume) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_VolumeResources retrieves all AWSEC2Instance_Volume items from a CloudFormation template
func GetAllAWSEC2Instance_Volume(template *Template) map[string]*AWSEC2Instance_Volume {

	results := map[string]*AWSEC2Instance_Volume{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_Volume{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_VolumeWithName retrieves all AWSEC2Instance_Volume items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_Volume(name string, template *Template) (*AWSEC2Instance_Volume, error) {

	result := &AWSEC2Instance_Volume{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_Volume{}, errors.New("resource not found")

}
