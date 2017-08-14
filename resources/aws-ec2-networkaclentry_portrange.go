package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::NetworkAclEntry.PortRange AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html
type AWSEC2NetworkAclEntry_PortRange struct {

	// From AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html#cfn-ec2-networkaclentry-portrange-from

	From int64 `json:"From"`

	// To AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html#cfn-ec2-networkaclentry-portrange-to

	To int64 `json:"To"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAclEntry_PortRange) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAclEntry.PortRange"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkAclEntry_PortRange) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkAclEntry_PortRangeResources retrieves all AWSEC2NetworkAclEntry_PortRange items from a CloudFormation template
func GetAllAWSEC2NetworkAclEntry_PortRange(template *Template) map[string]*AWSEC2NetworkAclEntry_PortRange {

	results := map[string]*AWSEC2NetworkAclEntry_PortRange{}
	for name, resource := range template.Resources {
		result := &AWSEC2NetworkAclEntry_PortRange{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkAclEntry_PortRangeWithName retrieves all AWSEC2NetworkAclEntry_PortRange items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2NetworkAclEntry_PortRange(name string, template *Template) (*AWSEC2NetworkAclEntry_PortRange, error) {

	result := &AWSEC2NetworkAclEntry_PortRange{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkAclEntry_PortRange{}, errors.New("resource not found")

}
