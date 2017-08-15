package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::NetworkInterface.PrivateIpAddressSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type AWSEC2NetworkInterface_PrivateIpAddressSpecification struct {

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
func (r *AWSEC2NetworkInterface_PrivateIpAddressSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterface.PrivateIpAddressSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkInterface_PrivateIpAddressSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NetworkInterface_PrivateIpAddressSpecificationResources retrieves all AWSEC2NetworkInterface_PrivateIpAddressSpecification items from a CloudFormation template
func GetAllAWSEC2NetworkInterface_PrivateIpAddressSpecification(template *Template) map[string]*AWSEC2NetworkInterface_PrivateIpAddressSpecification {

	results := map[string]*AWSEC2NetworkInterface_PrivateIpAddressSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEC2NetworkInterface_PrivateIpAddressSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NetworkInterface_PrivateIpAddressSpecificationWithName retrieves all AWSEC2NetworkInterface_PrivateIpAddressSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2NetworkInterface_PrivateIpAddressSpecification(name string, template *Template) (*AWSEC2NetworkInterface_PrivateIpAddressSpecification, error) {

	result := &AWSEC2NetworkInterface_PrivateIpAddressSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NetworkInterface_PrivateIpAddressSpecification{}, errors.New("resource not found")

}
