package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.Ebs AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html
type AWSEC2Instance_Ebs struct {

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-deleteontermination
	DeleteOnTermination bool `json:"DeleteOnTermination"`

	// Encrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-encrypted
	Encrypted bool `json:"Encrypted"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-iops
	Iops int64 `json:"Iops"`

	// SnapshotId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-snapshotid
	SnapshotId string `json:"SnapshotId"`

	// VolumeSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-volumesize
	VolumeSize int64 `json:"VolumeSize"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html#cfn-ec2-blockdev-template-volumetype
	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_Ebs) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.Ebs"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_Ebs) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_EbsResources retrieves all AWSEC2Instance_Ebs items from a CloudFormation template
func GetAllAWSEC2Instance_Ebs(template *Template) map[string]*AWSEC2Instance_Ebs {

	results := map[string]*AWSEC2Instance_Ebs{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_Ebs{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_EbsWithName retrieves all AWSEC2Instance_Ebs items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_Ebs(name string, template *Template) (*AWSEC2Instance_Ebs, error) {

	result := &AWSEC2Instance_Ebs{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_Ebs{}, errors.New("resource not found")

}
