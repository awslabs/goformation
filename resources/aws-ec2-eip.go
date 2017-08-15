package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::EIP AWS CloudFormation Resource
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
func GetAllAWSEC2EIP(template *Template) map[string]*AWSEC2EIP {

	results := map[string]*AWSEC2EIP{}
	for name, resource := range template.Resources {
		result := &AWSEC2EIP{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2EIPWithName retrieves all AWSEC2EIP items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2EIP(name string, template *Template) (*AWSEC2EIP, error) {

	result := &AWSEC2EIP{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2EIP{}, errors.New("resource not found")

}
