package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPNGateway AWS CloudFormation Resource (AWS::EC2::VPNGateway)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html
type AWSEC2VPNGateway struct {

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html#cfn-ec2-vpngateway-tags
	Tags []Tag `json:"Tags"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gateway.html#cfn-ec2-vpngateway-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNGateway) AWSCloudFormationType() string {
	return "AWS::EC2::VPNGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNGateway) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPNGatewayResources retrieves all AWSEC2VPNGateway items from a CloudFormation template
func GetAllAWSEC2VPNGatewayResources(template *Template) map[string]*AWSEC2VPNGateway {

	results := map[string]*AWSEC2VPNGateway{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPNGateway{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPNGatewayWithName retrieves all AWSEC2VPNGateway items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2VPNGatewayWithName(name string, template *Template) (*AWSEC2VPNGateway, error) {

	result := &AWSEC2VPNGateway{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPNGateway{}, errors.New("resource not found")

}
