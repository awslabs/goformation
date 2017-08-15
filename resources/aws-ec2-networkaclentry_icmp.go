package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::NetworkAclEntry.Icmp AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html
type AWSEC2NetworkAclEntry_Icmp struct {

	// Code AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-code

	Code int64 `json:"Code"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-type

	Type int64 `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAclEntry_Icmp) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAclEntry.Icmp"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkAclEntry_Icmp) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkAclEntry_IcmpResources retrieves all AWSEC2NetworkAclEntry_Icmp items from a CloudFormation template
func GetAllAWSEC2NetworkAclEntry_Icmp(template *Template) map[string]*AWSEC2NetworkAclEntry_Icmp {

	results := map[string]*AWSEC2NetworkAclEntry_Icmp{}
	for name, resource := range template.Resources {
		result := &AWSEC2NetworkAclEntry_Icmp{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkAclEntry_IcmpWithName retrieves all AWSEC2NetworkAclEntry_Icmp items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2NetworkAclEntry_Icmp(name string, template *Template) (*AWSEC2NetworkAclEntry_Icmp, error) {

	result := &AWSEC2NetworkAclEntry_Icmp{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkAclEntry_Icmp{}, errors.New("resource not found")

}
