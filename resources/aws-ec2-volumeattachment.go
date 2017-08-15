package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VolumeAttachment AWS CloudFormation Resource (AWS::EC2::VolumeAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
type AWSEC2VolumeAttachment struct {

	// Device AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-device
	Device string `json:"Device"`

	// InstanceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-instanceid
	InstanceId string `json:"InstanceId"`

	// VolumeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html#cfn-ec2-ebs-volumeattachment-volumeid
	VolumeId string `json:"VolumeId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VolumeAttachment) AWSCloudFormationType() string {
	return "AWS::EC2::VolumeAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VolumeAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VolumeAttachmentResources retrieves all AWSEC2VolumeAttachment items from a CloudFormation template
func GetAllAWSEC2VolumeAttachmentResources(template *Template) map[string]*AWSEC2VolumeAttachment {

	results := map[string]*AWSEC2VolumeAttachment{}
	for name, resource := range template.Resources {
		result := &AWSEC2VolumeAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VolumeAttachmentWithName retrieves all AWSEC2VolumeAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2VolumeAttachmentWithName(name string, template *Template) (*AWSEC2VolumeAttachment, error) {

	result := &AWSEC2VolumeAttachment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VolumeAttachment{}, errors.New("resource not found")

}
