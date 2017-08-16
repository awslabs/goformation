package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2EIP AWS CloudFormation Resource (AWS::EC2::EIP)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
type AWSEC2EIP struct {

	// Domain AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-domain
	Domain string `json:"Domain"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-instanceid
	InstanceId string `json:"InstanceId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EIP) AWSCloudFormationType() string {
	return "AWS::EC2::EIP"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2EIP) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2EIPResources retrieves all AWSEC2EIP items from a CloudFormation template
func (t *Template) GetAllAWSEC2EIPResources() map[string]*AWSEC2EIP {

	results := map[string]*AWSEC2EIP{}
	for name, resource := range t.Resources {
		result := &AWSEC2EIP{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2EIPWithName retrieves all AWSEC2EIP items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2EIPWithName(name string) (*AWSEC2EIP, error) {

	result := &AWSEC2EIP{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2EIP{}, errors.New("resource not found")

}
