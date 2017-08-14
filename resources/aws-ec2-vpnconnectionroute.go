package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::VPNConnectionRoute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html
type AWSEC2VPNConnectionRoute struct {

	// DestinationCidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-cidrblock

	DestinationCidrBlock string `json:"DestinationCidrBlock"`

	// VpnConnectionId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-connectionid

	VpnConnectionId string `json:"VpnConnectionId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNConnectionRoute) AWSCloudFormationType() string {
	return "AWS::EC2::VPNConnectionRoute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNConnectionRoute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPNConnectionRouteResources retrieves all AWSEC2VPNConnectionRoute items from a CloudFormation template
func GetAllAWSEC2VPNConnectionRoute(template *Template) map[string]*AWSEC2VPNConnectionRoute {

	results := map[string]*AWSEC2VPNConnectionRoute{}
	for name, resource := range template.Resources {
		result := &AWSEC2VPNConnectionRoute{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPNConnectionRouteWithName retrieves all AWSEC2VPNConnectionRoute items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2VPNConnectionRoute(name string, template *Template) (*AWSEC2VPNConnectionRoute, error) {

	result := &AWSEC2VPNConnectionRoute{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPNConnectionRoute{}, errors.New("resource not found")

}
