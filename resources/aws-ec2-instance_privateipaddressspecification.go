package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Instance.PrivateIpAddressSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type AWSEC2Instance_PrivateIpAddressSpecification struct {

	// Primary AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-primary
	Primary bool `json:"Primary"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-privateipaddress
	PrivateIpAddress string `json:"PrivateIpAddress"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_PrivateIpAddressSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.PrivateIpAddressSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_PrivateIpAddressSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2Instance_PrivateIpAddressSpecificationResources retrieves all AWSEC2Instance_PrivateIpAddressSpecification items from a CloudFormation template
func GetAllAWSEC2Instance_PrivateIpAddressSpecification(template *Template) map[string]*AWSEC2Instance_PrivateIpAddressSpecification {

	results := map[string]*AWSEC2Instance_PrivateIpAddressSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEC2Instance_PrivateIpAddressSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2Instance_PrivateIpAddressSpecificationWithName retrieves all AWSEC2Instance_PrivateIpAddressSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Instance_PrivateIpAddressSpecification(name string, template *Template) (*AWSEC2Instance_PrivateIpAddressSpecification, error) {

	result := &AWSEC2Instance_PrivateIpAddressSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Instance_PrivateIpAddressSpecification{}, errors.New("resource not found")

}
