package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPNGatewayRoutePropagation AWS CloudFormation Resource (AWS::EC2::VPNGatewayRoutePropagation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
type AWSEC2VPNGatewayRoutePropagation struct {

	// RouteTableIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-routetableids
	RouteTableIds []string `json:"RouteTableIds"`

	// VpnGatewayId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-vpngatewayid
	VpnGatewayId string `json:"VpnGatewayId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNGatewayRoutePropagation) AWSCloudFormationType() string {
	return "AWS::EC2::VPNGatewayRoutePropagation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNGatewayRoutePropagation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPNGatewayRoutePropagationResources retrieves all AWSEC2VPNGatewayRoutePropagation items from a CloudFormation template
func (t *Template) GetAllAWSEC2VPNGatewayRoutePropagationResources() map[string]*AWSEC2VPNGatewayRoutePropagation {

	results := map[string]*AWSEC2VPNGatewayRoutePropagation{}
	for name, resource := range t.Resources {
		result := &AWSEC2VPNGatewayRoutePropagation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPNGatewayRoutePropagationWithName retrieves all AWSEC2VPNGatewayRoutePropagation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNGatewayRoutePropagationWithName(name string) (*AWSEC2VPNGatewayRoutePropagation, error) {

	result := &AWSEC2VPNGatewayRoutePropagation{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPNGatewayRoutePropagation{}, errors.New("resource not found")

}
